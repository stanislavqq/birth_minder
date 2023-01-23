package notify

import (
	bevent "BMinder/internal/model/bevent"
	"fmt"
	"github.com/rs/zerolog"
	"math"
	"strings"
	"time"
)

type FindNotifyJob struct {
	logger        zerolog.Logger
	eventRep      *bevent.Repository
	notifyTimes   []time.Duration
	messageFormat string
	NotifyChan    chan Notify
	debug         bool
}

func NewJob(repository *bevent.Repository, messageFormat string, notifyTimes []time.Duration, collector chan Notify, debug bool, logger zerolog.Logger) *FindNotifyJob {
	return &FindNotifyJob{
		logger:        logger,
		eventRep:      repository,
		NotifyChan:    collector,
		debug:         debug,
		notifyTimes:   notifyTimes,
		messageFormat: messageFormat,
	}
}

func (j *FindNotifyJob) Run() {

	for _, interval := range j.notifyTimes {
		eventList, err := j.findBirthEventByDuration(interval)
		if err != nil {
			j.logger.Err(err).Dur("interval event", interval).Msg("Не удалось получить список событий по промежутку времени")
		}

		if len(eventList) > 0 {
			if j.debug {
				j.logger.Debug().Dur("interval event", interval).Fields(eventList).Msg("Фигачим в работу события")
			}

			for _, event := range eventList {
				notify := makeNotify(event, interval, j.messageFormat)
				go notifyToChan(notify, j.NotifyChan)
			}
		}
		eventList = bevent.BirthEvents{}
	}
}

func parseFormatMessage(format string, params map[string]string) string {
	res := format
	for key, value := range params {
		res = strings.ReplaceAll(res, "{"+key+"}", value)
	}

	return res
}

func makeNotify(event bevent.BirthEvent, interval time.Duration, formatMessage string) Notify {
	afterTime := durationToStringFormat(interval)

	msg := parseFormatMessage(formatMessage, map[string]string{
		"fullname":  event.GetFullName(),
		"firstname": event.FirstName,
		"lastname":  event.LastName,
		"soon_time": afterTime,
	})
	//msg := fmt.Sprintf("🎉🎉🎉Напоминание: \n\n У %s - %s будет День рождения.", event.GetFullName(), afterTime)
	return NewNotify(msg, interval)
}

func notifyToChan(notify Notify, notifyPipe chan Notify) {
	notifyPipe <- notify
	time.Sleep(time.Second)
}

func durationToStringFormat(duration time.Duration) string {
	var res string
	switch duration {
	case time.Hour * 24:
		res = "завтра"
	case time.Hour * 24 * 2:
		res = "после завтра"
	case time.Hour * 24 * 7:
		res = "через неделю"
	default:
		//week := time.Hour * 24 * 7

		days := int(math.Round(duration.Hours() / 24))
		if weeks := days / 7; days%7 == 0 && weeks < 4 {
			res = fmt.Sprintf("через %d недели", weeks)
		} else {
			res = fmt.Sprintf("через %d дней", days)
		}

	}

	return res
}

func (j *FindNotifyJob) findBirthEventByDuration(duration time.Duration) (bevent.BirthEvents, error) {
	paramDay, paramMonth := j.getDayMonthFromDuration(duration)
	eventList, err := j.eventRep.GetListByDayMonth(paramDay, paramMonth)
	if err != nil {
		j.logger.Err(err).Msg("Не удалось получить список записей")
	}

	return eventList, err
}

func (j *FindNotifyJob) getDayMonthFromDuration(duration time.Duration) (int, int) {
	now := time.Now()
	timeTarget := now.Add(duration)

	return timeTarget.Day(), int(timeTarget.Month())
}
