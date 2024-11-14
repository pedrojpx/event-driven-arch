package createclient

import (
	"time"

	"github.com/pedrojpx/ms-wallet/internal/entity"
	"github.com/pedrojpx/ms-wallet/internal/gateway"
)

type CreateClientInputDTO struct {
	Name  string
	Email string
}

type CreateClientOutputDTO struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateClientUseCase struct {
	Gateway gateway.ClientGateway
}

func NewCreateClientUseCase(gw gateway.ClientGateway) *CreateClientUseCase {
	return &CreateClientUseCase{
		Gateway: gw,
	}
}

func (u *CreateClientUseCase) Execute(input CreateClientInputDTO) (*CreateClientOutputDTO, error) {
	c, err := entity.NewClient(input.Name, input.Email)
	if err != nil {
		return nil, err
	}
	err = u.Gateway.Save(c)
	if err != nil {
		return nil, err
	}

	return &CreateClientOutputDTO{
		ID:        c.ID,
		Name:      c.Name,
		Email:     c.Email,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}, nil
}
