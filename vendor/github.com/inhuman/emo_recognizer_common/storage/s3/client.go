package s3

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const (
	DefaultEndpoint  = "localhost:9000"
	DefaultAccessKey = "access-key"
	DefaultSecretKey = "secret-key"
	DefaultBucket    = "test-bucket"
	DefaultS3Secure  = true
)

var defaultConnectTimeout = time.Second * 2

type Option func(o *S3)

var ErrBucketNotExists = errors.New("you need use already created bucket")

type S3 struct {
	client       *minio.Client
	endpoint     string
	accessKey    string
	secretKey    string
	bucketName   string
	secure       bool
	createBucket bool
}

func NewS3(opts ...Option) (*S3, error) {
	storage := &S3{
		endpoint:   DefaultEndpoint,
		accessKey:  DefaultAccessKey,
		secretKey:  DefaultSecretKey,
		bucketName: DefaultBucket,
		secure:     DefaultS3Secure,
	}

	for _, o := range opts {
		o(storage)
	}

	if err := storage.connect(); err != nil {
		return nil, err
	}

	return storage, nil
}

func (s *S3) GetPublicURLByFileName(fileName string) string {
	return fmt.Sprintf("%s/%s/%s", s.client.EndpointURL().String(), s.bucketName, fileName)
}

func (s *S3) connect() error {
	minioClient, err := minio.New(s.endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(s.accessKey, s.secretKey, ""),
		Secure: s.secure,
	})
	if err != nil {
		return fmt.Errorf("failed create minio client: %w", err)
	}

	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), defaultConnectTimeout)
	defer cancel()

	b, err := minioClient.BucketExists(ctxWithTimeout, s.bucketName)
	if err != nil {
		return fmt.Errorf("failed check bucket exists: %w", err)
	}

	if !b && !s.createBucket {
		return ErrBucketNotExists
	}

	if !b && s.createBucket {
		err = minioClient.MakeBucket(ctxWithTimeout, s.bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return fmt.Errorf("error create bucket with name %s: %w", s.bucketName, err)
		}
	}

	s.client = minioClient

	return nil
}

func (s *S3) Write(ctx context.Context, filename string, size int64, r io.Reader) (err error) {
	_, err = s.client.PutObject(ctx, s.bucketName, filename, r, size,
		minio.PutObjectOptions{
			ContentType: "application/octet-stream",
		},
	)
	if err != nil {
		return fmt.Errorf("failed put object: %w", err)
	}

	return nil
}

func (s *S3) Read(ctx context.Context, fileName string) (io.Reader, error) {
	obj, err := s.client.GetObject(ctx, s.bucketName, fileName, minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed get object: %w", err)
	}
	defer obj.Close()

	b, err := io.ReadAll(obj)
	if err != nil {
		return nil, fmt.Errorf("failed read all body: %w", err)
	}

	return bytes.NewBuffer(b), nil
}
