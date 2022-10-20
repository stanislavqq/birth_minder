package repository

import (
	"database/sql"
	"fmt"
)

type AccountRepository struct {
	db *sql.DB
}

func Create(db *sql.DB) AccountRepository {
	return AccountRepository{db: db}
}

func (a *AccountRepository) GetAccount(ID int32) Account {
	var account Account
	if err := a.db.QueryRow("SELECT * FROM accounts WHERE id = ?", ID).Scan(&account.ID, &account.CreatedAt); err != nil {
		fmt.Errorf("error get Account %d", ID)
		return Account{}
	}

	return account
}
