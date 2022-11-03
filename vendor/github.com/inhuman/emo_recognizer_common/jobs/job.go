package jobs

import (
	"fmt"
	"time"
)

type JobStatus string

const (
	JobStatusPlanned                 JobStatus = "planned"
	JobStatusFileUploaded            JobStatus = "file_uploaded"
	JobStatusNoiseWrapStarted        JobStatus = "noise_wrap_started"
	JobStatusNoiseWrapComplete       JobStatus = "noise_wrap_complete"
	JobStatusNoiseWrapError          JobStatus = "noise_wrap_error"
	JobStatusSpeechRecognizeStarted  JobStatus = "speech_recognize_started"
	JobStatusSpeechRecognizeComplete JobStatus = "speech_recognize_complete"
	JobStatusSpeechRecognizeError    JobStatus = "speech_recognize_error"
	JobStatusComplete                JobStatus = "complete"
	JobStatusCancelled               JobStatus = "cancelled"
)

type Job struct {
	UUID           string
	Status         JobStatus
	Filename       string
	RecognizedText string
	Strategy       Strategy
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

const ext = "wav"

func (j *Job) OriginalFileName() string {
	return fmt.Sprintf("%s_original.%s", j.UUID, ext)
}

func (j *Job) CleanFileName() string {
	return fmt.Sprintf("%s_clean.%s", j.UUID, ext)
}

func StatusesToProcess() []JobStatus {
	return []JobStatus{
		JobStatusFileUploaded,
		JobStatusNoiseWrapComplete,
		JobStatusSpeechRecognizeComplete,
	}

}

type Strategy string
