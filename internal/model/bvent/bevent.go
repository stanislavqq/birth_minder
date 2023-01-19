package bevent

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

type Repository struct {
	db     *sqlx.DB
	logger zerolog.Logger
}

type BirthEvents []BirthEvent

type BirthEvent struct {
	ID        int32          `db:"id"`
	FirstName string         `db:"firstname"`
	LastName  string         `db:"lastname"`
	Day       int32          `db:"day"`
	Month     int32          `db:"month"`
	Year      sql.NullInt32  `db:"year"`
	Comment   sql.NullString `db:"comment"`
	AccountID int32          `db:"account_id"`
}

func NewRepository(db *sqlx.DB, logger zerolog.Logger) *Repository {
	return &Repository{
		db:     db,
		logger: logger,
	}
}

func (r *Repository) GetList() (BirthEvents, error) {
	birthEvent := BirthEvent{}
	rows, err := r.db.Query("SELECT id, firstname, lastname, day, month, year, comment FROM birth_list")
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

	return eventList, nil
}
