package bevent

import (
	"database/sql"
	"github.com/rs/zerolog"
	"strings"
)

type Repository struct {
	db        *sql.DB
	logger    zerolog.Logger
	tableName string
	fields    []string
}

func NewRepository(db *sql.DB, logger zerolog.Logger) *Repository {
	return &Repository{
		db:        db,
		logger:    logger,
		tableName: "birth_list",
		fields:    []string{"id", "firstname", "lastname", "day", "month", "year", "comment"},
	}
}

func (r *Repository) GetListByDayMonth(day int, month int) (BirthEvents, error) {
	var birthEvent BirthEvent
	sqlQuery := "SELECT " + strings.Join(r.fields, ", ") + " FROM " + r.tableName + " WHERE day = ? AND month = ?"

	rows, err := r.db.Query(sqlQuery, day, month)
	if err != nil {
		r.logger.Err(err).Msg("Не удалось выполнить запрос")
		return nil, err
	}
	eventList := BirthEvents{}
	for rows.Next() {
		err := rows.Scan(&birthEvent.ID, &birthEvent.FirstName, &birthEvent.LastName, &birthEvent.Day, &birthEvent.Month, &birthEvent.Year, &birthEvent.Comment)
		if err != nil {
			r.logger.Err(err).Msg("Не удалось выполнить scan")
			return nil, err
		}
		eventList = append(eventList, birthEvent)
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			r.logger.Error().Err(err).Msg("Не удалось закрыть соединение")
		}
	}(rows)

	if len(eventList) == 0 {
		r.logger.Info().Msg("Не найдено записей для работы")
	}

	return eventList, nil
}

func (r *Repository) GetUpcomingBDay(day int, month int) (BirthEvents, error) {
	var birthEvent BirthEvent
	sqlQuery := "SELECT * FROM birth_list WHERE month = ? AND day > ? ORDER BY month ASC, day ASC"

	rows, err := r.db.Query(sqlQuery, month, day)
	if err != nil {
		r.logger.Err(err).Msg("Не удалось выполнить запрос")
		return nil, err
	}
	eventList := BirthEvents{}

	for rows.Next() {
		err := rows.Scan(&birthEvent.ID, &birthEvent.FirstName, &birthEvent.LastName, &birthEvent.Day, &birthEvent.Month, &birthEvent.Year, &birthEvent.Comment)
		if err != nil {
			r.logger.Err(err).Msg("Не удалось выполнить scan")
			return nil, err
		}
		eventList = append(eventList, birthEvent)
	}

	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			r.logger.Error().Err(err).Msg("Не удалось закрыть соединение")
		}
	}(rows)

	return eventList, nil
}
