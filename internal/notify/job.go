package notify

import (
	bevent "BMinder/internal/model/bvent"
	"fmt"
	"github.com/rs/zerolog"
	"time"
)

type NotifyJob struct {
	logger   zerolog.Logger
	eventRep bevent.Repository
}

func (j *NotifyJob) Run() {

	now := time.Now()
	duration, err := time.ParseDuration("7d")
	if err != nil {
		return
	}
	searchDate := now.Add(duration)

	fmt.Println(searchDate.Format("DD.MM.YY"))

	eventList, err := j.eventRep.GetList()
	if err != nil {
		j.logger.Err(err).Msg("Не удалось получить список записей")
	}

	fmt.Println(eventList)
}
