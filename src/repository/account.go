package repository

type Account struct {
	ID        int32  `db:"id"`
	CreatedAt string `db:"created_at"`
}
