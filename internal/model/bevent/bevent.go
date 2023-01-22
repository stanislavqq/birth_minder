package bevent

import (
	"database/sql"
)

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

func (e *BirthEvent) GetFullName() string {
	return e.FirstName + " " + e.LastName
}
