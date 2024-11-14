package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ID        string
	Name      string
	Email     string
	Accounts  []*Account
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewClient(name string, email string) (*Client, error) {
	c := &Client{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := c.Validate(); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Client) Validate() error {
	if c.Name == "" {
		return errors.New("name is required")
	}
	if c.Email == "" {
		return errors.New("email is required")
	}

	return nil
}

func (c *Client) Update(n, email string) error {
	c.Name = n
	c.Email = email
	c.UpdatedAt = time.Now()
	if err := c.Validate(); err != nil {
		return err
	}
	return nil
}

func (c *Client) AddAccount(acc *Account) error {
	if acc.Client.ID != c.ID {
		return errors.New("account does not belong to client")
	}
	c.Accounts = append(c.Accounts, acc)
	return nil
}
