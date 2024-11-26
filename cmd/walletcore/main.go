package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pedrojpx/ms-wallet/internal/database"
	"github.com/pedrojpx/ms-wallet/internal/event"
	createaccount "github.com/pedrojpx/ms-wallet/internal/usecase/create_account"
	createclient "github.com/pedrojpx/ms-wallet/internal/usecase/create_client"
	createtransaction "github.com/pedrojpx/ms-wallet/internal/usecase/create_transaction"
	"github.com/pedrojpx/ms-wallet/internal/web"
	"github.com/pedrojpx/ms-wallet/internal/web/webserver"
	"github.com/pedrojpx/ms-wallet/pkg/events"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "127.0.0.1", "3306", "wallet"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	eventDispatcher := events.NewEventDispatcher()
	// eventDispatcher.Register("TrasactionCreated", handler)

	clientDB := database.NewClientDB(db)
	accountDB := database.NewAccountDB(db)
	transactionDB := database.NewTransactionDB(db)

	createClientUseCase := createclient.NewCreateClientUseCase(clientDB)
	createAccountUseCase := createaccount.NewCreateAccountUseCase(accountDB, clientDB)
	createTransactionUsecase := createtransaction.NewCreateTransactionUseCase(accountDB, transactionDB, eventDispatcher, event.NewTransactionCreatedEvent())

	webserver := webserver.NewWebServer(":3000")
	createClientHandler := web.NewWebClientHandler(*createClientUseCase)
	createAccountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	createTransactionHandler := web.NewWebTransactionHandler(*createTransactionUsecase)

	webserver.AddHandler("/clients", createClientHandler.CreateClient)
	webserver.AddHandler("/accounts", createAccountHandler.CreateAccount)
	webserver.AddHandler("/transactions", createTransactionHandler.CreateTransaction)
	webserver.Start()
}
