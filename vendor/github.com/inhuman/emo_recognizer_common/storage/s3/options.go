package s3

import "github.com/minio/minio-go/v7"

func WithMinioClient(client *minio.Client) Option {
	return func(o *S3) {
		o.client = client
	}
}

func WithAccessKey(accessKey string) Option {
	return func(o *S3) {
		o.accessKey = accessKey
	}
}

func WithSecretKey(secretKey string) Option {
	return func(o *S3) {
		o.secretKey = secretKey
	}
}

func WithBucketName(bucketName string) Option {
	return func(o *S3) {
		o.bucketName = bucketName
	}
}

func WithEndpoint(endpoint string) Option {
	return func(o *S3) {
		o.endpoint = endpoint
	}
}

func WithSecure(secure bool) Option {
	return func(o *S3) {
		o.secure = secure
	}
}

func WithCreateBucket(create bool) Option {
	return func(o *S3) {
		o.createBucket = create
	}
}
