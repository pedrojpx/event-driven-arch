package balanceServices

import (
	"database/sql"
	"encoding/json"
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/pedrojpx/ms-wallet/pkg/kafka"
)

type BalanceKafkaConsumer struct {
	db       *sql.DB
	consumer *kafka.Consumer
}

type BalanceConsumerInputDTO struct {
	Name    string `json:"Name"`
	Payload struct {
		AccountIDFrom  string  `json:"account_from"`
		BalanceAccFrom float64 `json:"balance_account_from"`
		AccountIDTo    string  `json:"account_to"`
		BalanceAccTo   float64 `json:"balance_account_to"`
	} `json:"Payload"`
	Date string `json:"Date"`
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
			fmt.Println(string((<-m).Value))
			fmt.Println(received)
			b.saveBalances(received)
			fmt.Println("received ====")
		}
	}(msg)
}

func (b *BalanceKafkaConsumer) saveBalances(balances BalanceConsumerInputDTO) error {
	query := "insert into accounts (id, balance) values (?, ?) ON DUPLICATE KEY UPDATE balance=?"

	fmt.Println("inserting into db")
	fmt.Println(balances)

	stmt, err := b.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	//TODO: transformar em unit of work
	_, err = stmt.Exec(balances.Payload.AccountIDFrom, balances.Payload.BalanceAccFrom, balances.Payload.BalanceAccFrom)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(balances.Payload.AccountIDTo, balances.Payload.BalanceAccTo, balances.Payload.BalanceAccTo)
	if err != nil {
		return err
	}

	return nil
}
