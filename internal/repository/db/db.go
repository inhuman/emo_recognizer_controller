package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/inhuman/emo_recognizer_common/jobs"
	libpgx "github.com/inhuman/emo_recognizer_common/pgx"
	"github.com/inhuman/emo_recognizer_controller/internal/repository"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
)

const (
	defaultJobsLimit = 10
)

type Repository struct {
	db     libpgx.Pool
	logger *zap.Logger
}

func NewRepository(db libpgx.Pool, logger *zap.Logger) *Repository {
	return &Repository{
		db:     db,
		logger: logger,
	}
}

func (r *Repository) GetJobs(ctx context.Context, filter repository.GetJobsFilter) ([]*jobs.Job, error) {
	queryString, args, err := getQueryGetJobs(filter)
	if err != nil {
		return nil, fmt.Errorf("failed to generate sql for fetch jobs by filter: %w", err)
	}

	rows, err := r.db.Query(ctx, queryString, args...)
	defer rows.Close()

	var jobsFromDB []*jobs.Job

	for rows.Next() {
		reportFromDB, scanErr := scanJob(rows)
		if scanErr != nil {
			return nil, fmt.Errorf("error scan job from db: %w", scanErr)
		}

		jobsFromDB = append(jobsFromDB, reportFromDB)
	}

	return jobsFromDB, nil
}

func getQueryGetJobs(filter repository.GetJobsFilter) (queryString string, args []interface{}, err error) {
	query := sq.Select("uuid", "status", "file_name", "strategy", "recognized_text", "created_at", "updated_at").
		From("jobs")

	query = query.PlaceholderFormat(sq.Dollar)

	query = applyFilterToQuery(query, filter)

	query = query.OrderBy("id DESC")

	sqlArgs, sqlStr, err := query.ToSql()
	if err != nil {
		return "", nil, fmt.Errorf("error building query string: %w", err)
	}

	return sqlArgs, sqlStr, nil
}

func applyFilterToQuery(query sq.SelectBuilder, filter repository.GetJobsFilter) sq.SelectBuilder {
	if filter.Status != "" {
		query = query.Where(sq.Eq{"status": filter.Status})
	}

	if filter.Strategy != "" {
		query = query.Where(sq.Eq{"strategy": filter.Strategy})
	}

	if filter.Offset != 0 {
		query = query.Offset(uint64(filter.Offset))
	}

	switch {
	case filter.Limit == 0:
		query = query.Limit(defaultJobsLimit)
	case filter.Limit > 0:
		query = query.Limit(uint64(filter.Limit))
	}

	return query
}

func scanJob(row libpgx.Scanner) (*jobs.Job, error) {
	jobFromDB := jobs.Job{}

	var recognizedText sql.NullString

	err := row.Scan(
		&jobFromDB.UUID, &jobFromDB.Status, &jobFromDB.Filename, &jobFromDB.Strategy, &recognizedText,
		&jobFromDB.CreatedAt, &jobFromDB.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("error fetch jobs from db: %w", err)
	}

	if recognizedText.Valid {
		jobFromDB.RecognizedText = recognizedText.String
	}

	return &jobFromDB, nil
}

const queryGetJobByUUID = `
SELECT "uuid", "status", "file_name", "strategy", "recognized_text", "created_at", "updated_at"
FROM jobs
WHERE uuid = $1;
`

func (r *Repository) GetJobByUUID(ctx context.Context, jobUUID string) (*jobs.Job, error) {
	row := r.db.QueryRow(ctx, queryGetJobByUUID, jobUUID)
	jobFromDb, err := scanJob(row)
	if err != nil {
		return nil, err
	}

	return jobFromDb, nil
}

const queryGetJobToProcess = `
SELECT "uuid", "status", "file_name", "strategy", "recognized_text", "created_at", "updated_at"
FROM jobs
WHERE "status" IN (%s)
ORDER BY created_at
LIMIT 1
`

func (r *Repository) GetJobToProcess(ctx context.Context) (*jobs.Job, error) {
	var inClause string

	statuses := jobs.StatusesToProcess()

	for i := range statuses {
		inClause += fmt.Sprintf(`'%s',`, statuses[i])
	}

	inClause = strings.TrimRight(inClause, ",")

	row := r.db.QueryRow(ctx, fmt.Sprintf(queryGetJobToProcess, inClause))

	jobFromDb, err := scanJob(row)
	if err != nil {
		return nil, err
	}
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf("failed to fetch job: %w", err)
	}

	return jobFromDb, nil
}

const queryCreateJob = `
INSERT INTO jobs ("file_name") VALUES ($1) RETURNING "uuid"
`

func (r *Repository) CreateJob(ctx context.Context, jobToCreate *jobs.Job) error {
	err := r.db.QueryRow(ctx, queryCreateJob, jobToCreate.Filename).Scan(&jobToCreate.UUID)
	if err != nil {
		return err
	}

	return nil
}

const queryUpdateStatusByUUID = `
UPDATE jobs SET "status" = $1
WHERE "uuid" = $2
`

func (r *Repository) UpdateStatusByUUID(ctx context.Context, jobUUID string, status jobs.JobStatus) error {
	_, err := r.db.Exec(ctx, queryUpdateStatusByUUID, status, jobUUID)
	if err != nil {
		return err
	}

	return nil
}

const queryUpdateRecognizedText = `
UPDATE jobs SET "recognized_text" = $1
WHERE "uuid" = $2
`

func (r *Repository) UpdateRecognizedText(ctx context.Context, jobUUID string, text string) error {
	_, err := r.db.Exec(ctx, queryUpdateRecognizedText, text, jobUUID)
	if err != nil {
		return err
	}

	return nil
}
