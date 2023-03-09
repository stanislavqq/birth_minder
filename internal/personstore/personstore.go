package personstore

import (
	"BMinder/internal/model/bevent"
	"database/sql"
	"github.com/rs/zerolog"
)

type PersonStore struct {
	db     *sql.DB
	logger zerolog.Logger
}

func New(db *sql.DB, logger zerolog.Logger) *PersonStore {
	return &PersonStore{
		db:     db,
		logger: logger,
	}
}

func (s *PersonStore) CreatePerson() int {
	return 0
}

func (s *PersonStore) GetPerson() bevent.BirthEvent {
	return bevent.BirthEvent{}
}

func (s *PersonStore) DeletePerson() error {
	return nil
}

func (s *PersonStore) GetPersonAll() []bevent.BirthEvent {
	rows, err := s.db.Query("SELECT * FROM birth_list ORDER BY month, day")
	if err != nil {
		s.logger.Error().Err(err).Msg("Не удалось выполнить запрос на получение списка")
	}

	var result []bevent.BirthEvent

	for rows.Next() {
		var person bevent.BirthEvent
		if err := rows.Scan(&person.ID, &person.FirstName, &person.LastName, &person.Day, &person.Month, &person.Year, &person.Comment); err != nil {
			s.logger.Error().Err(err).Msg("Не удалось конвертировать ")
		}

		result = append(result, person)
	}

	return result
}
