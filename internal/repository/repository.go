package repository

import (
	"context"

	"github.com/inhuman/emo_recognizer_common/jobs"
)

type GetJobsFilter struct {
	Status   jobs.JobStatus
	Strategy jobs.Strategy
	Limit    int
	Offset   int
}
type Repository interface {
	GetJobs(ctx context.Context, filter GetJobsFilter) ([]*jobs.Job, error)
	GetJobByUUID(ctx context.Context, jobUUID string) (*jobs.Job, error)
	GetJobToProcess(ctx context.Context) (*jobs.Job, error)
	CreateJob(ctx context.Context, jobToCreate *jobs.Job) error
	UpdateStatusByUUID(ctx context.Context, jobUUID string, status jobs.JobStatus) error
	UpdateRecognizedText(ctx context.Context, jobUUID string, text string) error
}
