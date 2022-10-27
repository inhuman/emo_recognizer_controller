package jobs

import "time"

type JobStatus string

const (
	JobStatusPlanned                 JobStatus = "planned"
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
	UUID      string
	Status    JobStatus
	Filename  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
