package createaccount

import (
	"github.com/pedrojpx/ms-wallet/internal/entity"
	"github.com/pedrojpx/ms-wallet/internal/gateway"
)

type CreateAccountInputDTO struct {
	ClientID string `json:"client_id"`
}

type CreateAccountOutputDTO struct {
	ID string `json:"account_id"`
}

type CreateAccountUseCase struct {
	AccGW gateway.AccountGateway
	CliGw gateway.ClientGateway
}

func NewCreateAccountUseCase(a gateway.AccountGateway, c gateway.ClientGateway) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		AccGW: a,
		CliGw: c,
	}
}

func (uc *CreateAccountUseCase) Execute(input CreateAccountInputDTO) (*CreateAccountOutputDTO, error) {
	c, err := uc.CliGw.FindByID(input.ClientID)
	if err != nil {
		return nil, err
	}

	acc := entity.NewAccount(c)
	if err := uc.AccGW.Save(acc); err != nil {
		return nil, err
	}

	return &CreateAccountOutputDTO{ID: acc.ID}, nil
}
