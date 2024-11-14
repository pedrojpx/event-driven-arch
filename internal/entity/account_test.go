package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	c, _ := NewClient("john", "j@j")
	acc := NewAccount(c)

	assert.NotNil(t, acc)
	assert.Equal(t, c.ID, acc.Client.ID)
}

func TestCreateAccountWithNilClient(t *testing.T) {
	acc := NewAccount(nil)
	assert.Nil(t, acc)
}

func TestCreditAccount(t *testing.T) {
	c, _ := NewClient("John", "j@j")
	acc := NewAccount(c)
	acc.Credit(10)

	assert.Equal(t, acc.Balance, float64(10))
}

func TestDebitAccount(t *testing.T) {
	c, _ := NewClient("John", "j@j")
	acc := NewAccount(c)
	acc.Credit(10)
	acc.Debit(1)

	assert.Equal(t, acc.Balance, float64(9))
}
