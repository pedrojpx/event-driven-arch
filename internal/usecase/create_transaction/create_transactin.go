package createtransaction

import (
	"github.com/pedrojpx/ms-wallet/internal/entity"
	"github.com/pedrojpx/ms-wallet/internal/gateway"
)

type CreateTrasactionInputDTO struct {
	AccountIDFrom string
	AccountIDTo   string
	Amount        float64
}

type CreateAccountOutputDTO struct {
	TransactionID string
}

type CreateTransactionUsecase struct {
	accGw gateway.AccountGateway
	traGw gateway.TransactionGateway
}

func NewCreateTransactionUseCase(a gateway.AccountGateway, t gateway.TransactionGateway) *CreateTransactionUsecase {
	return &CreateTransactionUsecase{
		accGw: a,
		traGw: t,
	}
}

func (uc *CreateTransactionUsecase) Execute(input CreateTrasactionInputDTO) (*CreateAccountOutputDTO, error) {
	from, err := uc.accGw.FindByID(input.AccountIDFrom)
	if err != nil {
		return nil, err
	}
	to, err := uc.accGw.FindByID(input.AccountIDTo)
	if err != nil {
		return nil, err
	}
	transactino, err := entity.NewTransaction(from, to, input.Amount)
	if err != nil {
		return nil, err
	}
	err = uc.traGw.Create(transactino)
	if err != nil {
		return nil, err
	}
	return &CreateAccountOutputDTO{TransactionID: transactino.ID}, nil
}
