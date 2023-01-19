package notify

import (
	"BMinder/internal/model/bvent"
	"github.com/rs/zerolog"
)

func NewJob(repository *bevent.Repository, logger zerolog.Logger) *NotifyJob {
	return &NotifyJob{
		logger:   logger,
		eventRep: *repository,
	}
}
