package repository

import "database/sql"

type User struct {
	ID        int32          `db:"id"`
	Name      string         `db:"name"`
	Day       int32          `db:"day"`
	Month     int32          `db:"month"`
	Year      sql.NullInt32  `db:"year"`
	Comment   sql.NullString `db:"comment"`
	AccountID int32          `db:"account_id"`
}
