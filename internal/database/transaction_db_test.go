package database

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pedrojpx/ms-wallet/internal/entity"
	"github.com/stretchr/testify/suite"
)

type TransactinoDBTestSuite struct {
	suite.Suite
	db    *sql.DB
	c1    *entity.Client
	c2    *entity.Client
	aFrom *entity.Account
	aTo   *entity.Account
	tdb   *TransactinoDB
}

func (s *TransactinoDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	_, err = db.Exec("Create table clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	s.Nil(err)
	_, err = db.Exec("Create table accounts (id varchar(255), client_id varchar(255), balance int, created_at date)")
	s.Nil(err)
	_, err = db.Exec("Create table transactions (id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount int, created_at date)")
	s.Nil(err)
	s.tdb = NewTransactionDB(db)

	client, err := entity.NewClient("a", "b")
	s.Nil(err)
	client2, err := entity.NewClient("c", "d")
	s.Nil(err)
	s.c1 = client
	s.c2 = client2

	from := entity.NewAccount(s.c1)
	from.Balance = 1000
	s.aFrom = from

	to := entity.NewAccount(s.c2)
	to.Balance = 1000
	s.aTo = to
}

func (s *TransactinoDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE transactions")
}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(TransactinoDBTestSuite))
}

func (s *TransactinoDBTestSuite) TestCreate() {
	t, err := entity.NewTransaction(s.aFrom, s.aTo, 100)
	s.Nil(err)
	err = s.tdb.Create(t)
	s.Nil(err)
}
