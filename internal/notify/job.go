package notify

import (
	bevent "BMinder/internal/model/bvent"
	"fmt"
	"github.com/rs/zerolog"
	"time"
)

type Job struct {
	logger   zerolog.Logger
	eventRep *bevent.Repository
}

func (j *Job) Run() {
	duration := (7 * 24) * time.Hour

	rules := bevent.Rules{
		bevent.NewTimeRule(duration),
	}

	eventList, err := j.eventRep.GetListWithRule(rules)
	if err != nil {
		j.logger.Err(err).Msg("Не удалось получить список записей")
	}

	fmt.Println(eventList)
}
