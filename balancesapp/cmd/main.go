package main

import (
	"database/sql"
	"fmt"
	"net/http"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pedrojpx/ms-wallet/balancesapp/balanceServices"
	"github.com/pedrojpx/ms-wallet/pkg/kafka"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql2", "3306", "balances"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	service := balanceServices.NewBalanceService(db)

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
	balanceServices.NewKafkaConsumer(consumer, db).Start()

	fmt.Println("balances app is running")
	http.ListenAndServe(port, router)
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("balances app is up"))
}
