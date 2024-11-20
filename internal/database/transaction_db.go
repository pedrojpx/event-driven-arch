package database

import (
	"database/sql"

	"github.com/pedrojpx/ms-wallet/internal/entity"
)

type TransactinoDB struct {
	DB *sql.DB
}

func NewTransactionDB(db *sql.DB) *TransactinoDB {
	return &TransactinoDB{
		DB: db,
	}
}

func (t *TransactinoDB) Create(tr *entity.Transaction) error {
	stmt, err := t.DB.Prepare("INSERT INTO transactions (id, account_id_from, account_id_to, amount, created_at) values (?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(tr.ID, tr.From.ID, tr.To.ID, tr.Amount, tr.CreatedAt)
	return err
}
