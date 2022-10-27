package repository

import "github.com/inhuman/emo_recognizer_common/jobs"

type Repository interface {
}

type GetJobsFilter struct {
	Status jobs.JobStatus
	Limit  int
	Offset int
}
