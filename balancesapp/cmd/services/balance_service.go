package BalanceService

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type BalanceService struct {
	db *sql.DB
}

type BalanceResponse struct {
	AccId   string `json:"accId"`
	Balance int    `json:"balance"`
}

func BewBalanceService(db *sql.DB) *BalanceService {
	return &BalanceService{
		db: db,
	}
}

func (b *BalanceService) BalanceHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "accId")

	balance, err := b.findBalance(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	resp := &BalanceResponse{AccId: id, Balance: balance}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (b *BalanceService) findBalance(id string) (int, error) {
	var balance int
	stmt, err := b.db.Prepare("SELECT balance from accounts where id = ?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	if err := row.Scan(&balance); err != nil {
		return 0, err
	}
	return balance, nil
}
