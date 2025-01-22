package main

import (
	"context"
	"database/sql"
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pedrojpx/ms-wallet/internal/database"
	"github.com/pedrojpx/ms-wallet/internal/event"
	"github.com/pedrojpx/ms-wallet/internal/event/handler"
	createaccount "github.com/pedrojpx/ms-wallet/internal/usecase/create_account"
	createclient "github.com/pedrojpx/ms-wallet/internal/usecase/create_client"
	createtransaction "github.com/pedrojpx/ms-wallet/internal/usecase/create_transaction"
	"github.com/pedrojpx/ms-wallet/internal/web"
	"github.com/pedrojpx/ms-wallet/internal/web/webserver"
	"github.com/pedrojpx/ms-wallet/pkg/events"
	"github.com/pedrojpx/ms-wallet/pkg/kafka"
	"github.com/pedrojpx/ms-wallet/pkg/uow"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql", "3306", "wallet"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
	}
	kafkaProducer := kafka.NewKafkaProducer(&configMap)

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("TransactionCreated", handler.NewTransactionCreatedKafkaHandler(kafkaProducer))
	eventDispatcher.Register("BalanceUpdated", handler.NewUpdateBalanceKafkaHandler(kafkaProducer))

	clientDB := database.NewClientDB(db)
	accountDB := database.NewAccountDB(db)

	ctx := context.Background()
	uow := uow.NewUow(ctx, db)
	uow.Register("AccountDB", func(tx *sql.Tx) interface{} {
		return database.NewAccountDB(db)
	})
	uow.Register("TransactionDB", func(tx *sql.Tx) interface{} {
		return database.NewTransactionDB(db)
	})

	createClientUseCase := createclient.NewCreateClientUseCase(clientDB)
	createAccountUseCase := createaccount.NewCreateAccountUseCase(accountDB, clientDB)
	createTransactionUsecase := createtransaction.NewCreateTransactionUseCase(uow, eventDispatcher, event.NewTransactionCreatedEvent(), event.NewBalanceUpdatedEvent())

	webserver := webserver.NewWebServer(":8080")
	createClientHandler := web.NewWebClientHandler(*createClientUseCase)
	createAccountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	createTransactionHandler := web.NewWebTransactionHandler(*createTransactionUsecase)

	webserver.AddHandler("/clients", createClientHandler.CreateClient)
	webserver.AddHandler("/accounts", createAccountHandler.CreateAccount)
	webserver.AddHandler("/transactions", createTransactionHandler.CreateTransaction)

	fmt.Println("Server is running")
	webserver.Start()
}
