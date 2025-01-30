package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
	BalanceService "github.com/pedrojpx/ms-wallet/balancesapp/cmd/services"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql2", "3307", "balances"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	service := BalanceService.BewBalanceService(db)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/healthcheck", healthcheck)
	router.Get("/balances/{accId}", service.BalanceHandler)
	port := ":3003"

	fmt.Println("balances app is running")
	http.ListenAndServe(port, router)
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("balances app is up"))
}
