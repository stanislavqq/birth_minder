package notify

import (
	"BMinder/internal/model/bvent"
	"github.com/rs/zerolog"
)

func NewJob(repository *bevent.Repository, logger zerolog.Logger) *Job {
	return &Job{
		logger:   logger,
		eventRep: repository,
	}
}
