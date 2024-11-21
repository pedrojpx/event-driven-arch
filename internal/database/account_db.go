package database

import (
	"database/sql"

	"github.com/pedrojpx/ms-wallet/internal/entity"
)

type AccountDB struct {
	DB *sql.DB
}

func NewAccountDB(db *sql.DB) *AccountDB {
	return &AccountDB{
		DB: db,
	}
}

func (a *AccountDB) FindByID(id string) (*entity.Account, error) {
	var acc entity.Account
	var cl entity.Client
	acc.Client = &cl

	stmt, err := a.DB.Prepare("SELECT a.id, a.client_id, a.balance, a.created_at, c.id, c.name, c.email, c.created_at FROM accounts a INNER JOIN clients c ON a.client_id = c.id WHERE a.id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(
		&acc.ID, &acc.Client.ID, &acc.Balance, &acc.CreatedAt,
		&cl.ID, &cl.Name, &cl.Email, &cl.CreatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &acc, nil
}

func (a *AccountDB) Save(acc *entity.Account) error {
	stmt, err := a.DB.Prepare(`
		INSERT INTO 
			accounts (id, client_id, balance, created_at) VALUES (?,?,?,?)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(acc.ID, acc.Client.ID, acc.Balance, acc.CreatedAt)
	// if err != nil {
	// 	return err
	// }
	// return nil
	return err
}
