package bevent

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"strings"
)

type Repository struct {
	db        *sqlx.DB
	logger    zerolog.Logger
	items     BirthEvents
	tableName string
	fields    []string
}

func NewRepository(db *sqlx.DB, logger zerolog.Logger) *Repository {
	return &Repository{
		db:        db,
		logger:    logger,
		tableName: "birth_list",
		fields:    []string{"id", "firstname", "lastname", "day", "month", "year", "comment"},
	}
}

func (r *Repository) GetListWithRule(rules Rules) (BirthEvents, error) {
	var sqlConditions []string
	var birthEvent BirthEvent
	var condParams []string

	for key, rule := range rules {
		sqlCond, params := r.buildQueryRule(rule)
		if key == 0 {
			sqlConditions = append(sqlConditions, "WHERE "+sqlCond)
		} else {
			sqlConditions = append(sqlConditions, "AND "+sqlCond)
		}

		condParams = append(condParams, params...)
	}

	condStr := strings.Join(sqlConditions, " ")
	sqlQuery := "SELECT " + strings.Join(r.fields, ", ") + " FROM " + r.tableName + " " + condStr
	r.logger.Info().Str("sqlQuery", sqlQuery).Strs("params", condParams).Msg("Запрос")

	rows, err := r.db.Query(sqlQuery)

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
		r.items = append(eventList, birthEvent)
	}

	if len(r.items) == 0 {
		r.logger.Info().Msg("Не найдено записей для работы")
	}

	return r.items, nil
}

func (r *Repository) buildQueryRule(rule Rule) (string, []string) {
	paramStrList := []string{
		fmt.Sprint(rule.day), fmt.Sprint(rule.month),
	}

	return "day = ? AND month = ?", paramStrList
}

func (r *Repository) GetList() (BirthEvents, error) {
	birthEvent := BirthEvent{}

	rows, err := r.db.Query("SELECT " + fmt.Sprintf("%s, ", r.fields) + " FROM " + r.tableName)
	if err != nil {
		r.logger.Err(err).Msg("Не удалось выполнить запрос")
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			r.logger.Err(err).Msg("Не удалось закрыть rows.Close()")
		}
	}(rows)

	eventList := BirthEvents{}
	for rows.Next() {
		err := rows.Scan(&birthEvent.ID, &birthEvent.FirstName, &birthEvent.LastName, &birthEvent.Day, &birthEvent.Month, &birthEvent.Year, &birthEvent.Comment)
		if err != nil {
			r.logger.Err(err).Msg("Не удалось выполнить scan")
			return nil, err
		}
		r.items = append(eventList, birthEvent)
	}

	if len(r.items) == 0 {
		r.logger.Info().Msg("Не найдено записей для работы")
	}

	return r.items, nil
}

func (r *Repository) GetCount() int {
	return len(r.items)
}
