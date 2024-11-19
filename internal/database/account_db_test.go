package database

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pedrojpx/ms-wallet/internal/entity"
	"github.com/stretchr/testify/suite"
)

type AccountDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	accountDB *AccountDB
	client    *entity.Client
}

func (s *AccountDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	_, err = db.Exec("Create table clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	s.Nil(err)
	_, err = db.Exec("Create table accounts (id varchar(255), client_id varchar(255), balance int, created_at date)")
	s.Nil(err)
	s.accountDB = NewAccountDB(db)
	s.client, _ = entity.NewClient("john", "j@j")
}

func (s *AccountDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE accounts")
}

func TestAccountDBSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}

func (s *AccountDBTestSuite) TestSave() {
	account := entity.NewAccount(s.client)
	err := s.accountDB.Save(*account)
	s.Nil(err)
}

func (s *AccountDBTestSuite) TestFindByID() {
	s.db.Exec("Insert into clients(id,name,email,created_at) values (?,?,?,?)", s.client.ID, s.client.Name, s.client.Email, s.client.CreatedAt)
	account := entity.NewAccount(s.client)
	err := s.accountDB.Save(*account)
	s.Nil(err)

	accFromDB, err := s.accountDB.FindByID(account.ID)
	s.Nil(err)
	s.Equal(account.ID, accFromDB.ID)
	s.Equal(account.Client.ID, accFromDB.Client.ID)
	s.Equal(account.Balance, accFromDB.Balance)
}
