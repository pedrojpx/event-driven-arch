package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
	BalanceService "github.com/pedrojpx/ms-wallet/balancesapp/cmd/services"
	"github.com/pedrojpx/ms-wallet/pkg/kafka"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql2", "3307", "balances"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	service := BalanceService.NewBalanceService(db)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/healthcheck", healthcheck)
	router.Get("/balances/{accId}", service.BalanceHandler)
	port := ":3003"

	kafkaConfig := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "balances",
	}
	consumer := kafka.NewConsumer(&kafkaConfig, []string{"balances"})
	NewKafkaConsumer(consumer, db).Start()

	fmt.Println("balances app is running")
	http.ListenAndServe(port, router)
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("balances app is up"))
}

type BalanceKafkaConsumer struct {
	db       *sql.DB
	consumer *kafka.Consumer
}

type BalanceConsumerInputDTO struct {
	AccountIDFrom  string  `json:"account_from"`
	BalanceAccFrom float64 `json:"balance_account_from"`
	AccountIDTo    string  `json:"account_to"`
	BalanceAccTo   float64 `json:"balance_account_to"`
}

func NewKafkaConsumer(c *kafka.Consumer, db *sql.DB) *BalanceKafkaConsumer {
	return &BalanceKafkaConsumer{
		consumer: c,
		db:       db,
	}
}

func (b *BalanceKafkaConsumer) Start() {
	msg := make(chan *ckafka.Message)
	go func(m chan *ckafka.Message) {
		for {
			fmt.Println("\n=====\nwaiting msg")
			go b.consumer.Consume(m)
			var received BalanceConsumerInputDTO
			json.Unmarshal((<-m).Value, &received)
			fmt.Println(received)
			b.saveBalances(received)
			fmt.Println("received ====\n")
		}
	}(msg)
}

func (b *BalanceKafkaConsumer) saveBalances(balances BalanceConsumerInputDTO) error {
	query := "insert into accounts (id, balance) values (?, ?) ON DUPLICATE KEY UPDATE balance=?"

	stmt, err := b.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	//TODO: transformar em unit of work
	_, err = stmt.Exec(balances.AccountIDFrom, balances.BalanceAccFrom, balances.BalanceAccFrom)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(balances.BalanceAccTo, balances.BalanceAccTo, balances.BalanceAccTo)
	if err != nil {
		return err
	}

	return nil
}
