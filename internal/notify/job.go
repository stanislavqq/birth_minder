package notify

import (
	bevent "BMinder/internal/model/bevent"
	"fmt"
	"github.com/rs/zerolog"
	"time"
)

const day = time.Hour * 24
const week = day * 7

type Job struct {
	logger     zerolog.Logger
	eventRep   *bevent.Repository
	NotifyChan chan Notify
	debug      bool
}

func NewJob(repository *bevent.Repository, collector chan Notify, debug bool, logger zerolog.Logger) *Job {
	return &Job{
		logger:     logger,
		eventRep:   repository,
		NotifyChan: collector,
		debug:      debug,
	}
}

func (j *Job) Run() {

	intervals := []time.Duration{day, week}

	for _, interval := range intervals {
		eventList, err := j.findBirthEventByDuration(interval)
		if err != nil {
			j.logger.Err(err).Dur("interval event", interval).Msg("Не удалось получить список событий по промежутку времени")
		}

		if len(eventList) > 0 {
			if j.debug {
				j.logger.Debug().Dur("interval event", interval).Fields(eventList).Msg("Фигачим в работу события")
			}
			go toChan(eventList, interval, j.NotifyChan)
		}
		eventList = bevent.BirthEvents{}
	}
}

func toChan(eventList bevent.BirthEvents, interval time.Duration, notifyPipe chan Notify) {
	for _, event := range eventList {
		afterTime := durationToStringFormat(interval)
		msg := fmt.Sprintf("🎉🎉🎉Напоминание: \n\n У %s - %s будет День рождения.", event.GetFullName(), afterTime)
		notif := NewNotify(msg, interval)
		notifyPipe <- notif

		time.Sleep(time.Second)
	}
}

func durationToStringFormat(duration time.Duration) string {
	var res string
	switch duration {
	case day:
		res = "завтра"
	case week:
		res = "через неделю"
	default:
		res = fmt.Sprintf("через %sч", duration.Hours())
	}

	return res
}

func (j *Job) findBirthEventByDuration(duration time.Duration) (bevent.BirthEvents, error) {
	paramDay, paramMonth := j.getDayMonthFromDuration(duration)
	eventList, err := j.eventRep.GetListByDayMonth(paramDay, paramMonth)
	if err != nil {
		j.logger.Err(err).Msg("Не удалось получить список записей")
	}

	return eventList, err
}

func (j *Job) getDayMonthFromDuration(duration time.Duration) (int, int) {
	now := time.Now()
	timeTarget := now.Add(duration)

	return timeTarget.Day(), int(timeTarget.Month())
}
