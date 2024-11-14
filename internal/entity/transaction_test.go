package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	c1, _ := NewClient("c1", "j@j")
	acc1 := NewAccount(c1)
	acc1.Credit(100)

	c2, _ := NewClient("c1", "j@j")
	acc2 := NewAccount(c2)
	acc2.Credit(100)

	tr, err := NewTransaction(acc1, acc2, 50)

	assert.Nil(t, err)
	assert.NotNil(t, tr)
	assert.Equal(t, acc2.Balance, 150.0)
}

func TestCreateTransactionWithInsuficentBalance(t *testing.T) {
	c1, _ := NewClient("c1", "j@j")
	acc1 := NewAccount(c1)
	acc1.Credit(100)

	c2, _ := NewClient("c1", "j@j")
	acc2 := NewAccount(c2)
	acc2.Credit(100)

	tr, err := NewTransaction(acc1, acc2, 500)

	assert.Nil(t, tr)
	assert.NotNil(t, err)
	assert.Equal(t, acc2.Balance, 100.0)
	assert.Equal(t, acc1.Balance, 100.0)
	assert.Error(t, err, "insufficient funds")

}
