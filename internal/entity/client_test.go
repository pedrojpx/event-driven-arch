package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("John Doe", "j@j.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, client.Name, "John Doe")
	assert.Equal(t, client.Email, "j@j.com")
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	c, _ := NewClient("John Doe", "j@j.com")
	err := c.Update("John Doe update", "j@g.com")

	assert.Nil(t, err)
	assert.Equal(t, c.Name, "John Doe update")
	assert.Equal(t, c.Email, "j@g.com")
}

func TestUpdateClientWithInvalidArgs(t *testing.T) {
	c, _ := NewClient("John Doe", "j@j.com")
	err := c.Update("", "j@g.com")

	assert.Error(t, err, "name is required")
}

func TestAddAccountToClient(t *testing.T) {
	c, _ := NewClient("John Doe", "j@j.com")
	acc := NewAccount(c)
	err := c.AddAccount(acc)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(c.Accounts))
}
