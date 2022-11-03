package db

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/inhuman/emo_recognizer_common/dockerpg"
	"github.com/inhuman/emo_recognizer_common/jobs"
	libpgx "github.com/inhuman/emo_recognizer_common/pgx"
	"github.com/inhuman/emo_recognizer_controller/internal/repository"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

var (
	pgxPool libpgx.Pool
	logger  = zap.NewExample()
	pgl     = dockerpg.NewPgLogger(logger)
	loc     *time.Location
)

func TestMain(m *testing.M) {
	var err error

	loc, err = time.LoadLocation("Europe/Moscow")
	if err != nil {
		panic(loc)
	}

	dPg, err := dockerpg.NewDockerPg(
		dockerpg.WithMigrationPath("file://./../../../migrations"),
		dockerpg.WithContainerName("dbrepo_test"),
		dockerpg.WithRandomPgPort(),
	)

	defer func() {
		err := dPg.Close()
		if err != nil {
			panic(err)
		}
	}()

	if err != nil {
		panic(err)
	}

	pgArgs := &libpgx.PgArgs{
		Host:     dPg.AppliedOptions().PgHost,
		Port:     dPg.AppliedOptions().PgPort,
		DB:       dPg.AppliedOptions().PgDb,
		User:     dPg.AppliedOptions().PgUser,
		Password: dPg.AppliedOptions().PgPassword,
		SslMode:  dPg.AppliedOptions().PgSslMode,
	}

	conn, err := pgArgs.String()
	if err != nil {
		panic(fmt.Sprintf("error build pg connections string: %s", err))
	}

	pgxConf, err := pgxpool.ParseConfig(conn)
	if err != nil {
		panic(fmt.Errorf("error parse pg connections string: %w", err))
	}

	pgxConf.ConnConfig.Logger = pgl

	pgxPool, err = pgxpool.ConnectConfig(context.Background(), pgxConf)
	if err != nil {
		panic(fmt.Errorf("error create pg connection pool: %w", err))
	}

	if err != nil {
		panic(err)
	}

	m.Run()
}

func TestRepository_CreateJob(t *testing.T) {
	contextWithTimeout, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	repo := NewRepository(pgxPool, logger)
	defer func() {
		cleanup(t, contextWithTimeout, "jobs")
	}()

	testJob := jobs.Job{
		Filename: "test_sound",
	}

	err := repo.CreateJob(contextWithTimeout, &testJob)
	assert.NoError(t, err)

	_, err = uuid.Parse(testJob.UUID)
	assert.NoError(t, err)
}

func TestRepository_GetJobs(t *testing.T) {
	contextWithTimeout, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	repo := NewRepository(pgxPool, logger)
	defer func() {
		cleanup(t, contextWithTimeout, "jobs")
	}()

	err := libpgx.QueryFromFile(pgxPool, "./../../../testdata/add_jobs.sql")
	assert.NoError(t, err)

	jobsFromDb, err := repo.GetJobs(contextWithTimeout, repository.GetJobsFilter{})
	assert.NoError(t, err)
	assert.Equal(t, 2, len(jobsFromDb))
}

func TestRepository_GetJobsFilterStatus(t *testing.T) {
	contextWithTimeout, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	repo := NewRepository(pgxPool, logger)
	defer func() {
		cleanup(t, contextWithTimeout, "jobs")
	}()

	err := libpgx.QueryFromFile(pgxPool, "./../../../testdata/add_jobs.sql")
	assert.NoError(t, err)

	jobsFromDb, err := repo.GetJobs(contextWithTimeout, repository.GetJobsFilter{
		Status: jobs.JobStatusPlanned,
	})
	assert.NoError(t, err)
	assert.Equal(t, 1, len(jobsFromDb))
}

func TestRepository_GetJobByUUID(t *testing.T) {
	contextWithTimeout, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	repo := NewRepository(pgxPool, logger)
	defer func() {
		cleanup(t, contextWithTimeout, "jobs")
	}()

	err := libpgx.QueryFromFile(pgxPool, "./../../../testdata/add_jobs.sql")
	assert.NoError(t, err)

	testJobUUID := "af0748b4-621b-422a-a968-86e54bfd9372"

	jobFromDb, err := repo.GetJobByUUID(contextWithTimeout, testJobUUID)
	assert.NoError(t, err)
	assert.Equal(t, jobs.JobStatusPlanned, jobFromDb.Status)
}

func TestRepository_UpdateStatusByUUID(t *testing.T) {
	contextWithTimeout, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	repo := NewRepository(pgxPool, logger)
	defer func() {
		cleanup(t, contextWithTimeout, "jobs")
	}()

	err := libpgx.QueryFromFile(pgxPool, "./../../../testdata/add_jobs.sql")
	assert.NoError(t, err)

	testJobUUID := "af0748b4-621b-422a-a968-86e54bfd9372"

	var status jobs.JobStatus

	err = pgxPool.QueryRow(contextWithTimeout, `SELECT "status" FROM jobs WHERE "uuid" = $1`, testJobUUID).Scan(&status)
	assert.NoError(t, err)
	assert.Equal(t, jobs.JobStatusPlanned, status)

	err = repo.UpdateStatusByUUID(contextWithTimeout, testJobUUID, jobs.JobStatusCancelled)
	assert.NoError(t, err)

	err = pgxPool.QueryRow(contextWithTimeout, `SELECT "status" FROM jobs WHERE "uuid" = $1`, testJobUUID).Scan(&status)
	assert.NoError(t, err)
	assert.Equal(t, jobs.JobStatusCancelled, status)
}

func TestRepository_GetJobToProcess(t *testing.T) {
	contextWithTimeout, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	repo := NewRepository(pgxPool, logger)
	defer func() {
		cleanup(t, contextWithTimeout, "jobs")
	}()

	err := libpgx.QueryFromFile(pgxPool, "./../../../testdata/add_jobs.sql")
	assert.NoError(t, err)

	jobFromDb, err := repo.GetJobToProcess(contextWithTimeout)
	assert.NoError(t, err)
	assert.Equal(t, "af0748b4-621b-422a-a968-86e54bfd9372", jobFromDb.UUID)
}

func cleanup(t *testing.T, ctx context.Context, tableName string) {
	_, err := pgxPool.Exec(ctx, fmt.Sprintf("TRUNCATE TABLE %s", tableName))
	if err != nil {
		t.Fatalf("failed truncate db: %s", err)
	}
}
