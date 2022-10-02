package repository

import (
	"database/sql"
	"fmt"
	"log"
)

type UserRepository struct {
	db *sql.DB
}

type User struct {
	ID        int32          `db:"id"`
	Name      string         `db:"name"`
	Day       int32          `db:"day"`
	Month     int32          `db:"month"`
	Year      sql.NullInt32  `db:"year"`
	Comment   sql.NullString `db:"comment"`
	AccountID int32          `db:"account_id"`
}

func (r UserRepository) GetUsers() map[int32]User {
	rows, err := r.db.Query("SELECT * FROM 'birth_users'")
	if err != nil {
		log.Panic(err)
	}
	//var users []BDUser
	list := make(map[int32]User)

	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Day, &user.Month, &user.Year, &user.AccountID, &user.Comment)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(user)
		list[user.ID] = user
		fmt.Printf("%d. %s\n", user.ID, user.Name)
	}

	rows.Close()
	defer r.db.Close()

	return list
}
