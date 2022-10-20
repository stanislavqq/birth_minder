package repository

import (
	"database/sql"
	"fmt"
	"log"
)

type UserRepository struct {
	db *sql.DB
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

func (u *UserRepository) FindBirthday(accountID int32, day string, month string) User {
	var buser User
	rows, err := u.db.Query("SELECT * FROM birth_users WHERE account_id = ? AND day = ? AND month = ?", accountID, day, month)
	if err != nil {
		log.Panic(err)
	}

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
}
