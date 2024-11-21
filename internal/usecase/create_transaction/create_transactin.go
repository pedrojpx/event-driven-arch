package createtransaction

import (
	"github.com/pedrojpx/ms-wallet/internal/entity"
	"github.com/pedrojpx/ms-wallet/internal/gateway"
	"github.com/pedrojpx/ms-wallet/pkg/events"
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
	accGw                   gateway.AccountGateway
	traGw                   gateway.TransactionGateway
	eventDistpatcher        events.EventDispatcherInterface
	transactionCreatedEvent events.EventInterface
}

func NewCreateTransactionUseCase(a gateway.AccountGateway, t gateway.TransactionGateway, ed events.EventDispatcherInterface, e events.EventInterface) *CreateTransactionUsecase {
	return &CreateTransactionUsecase{
		accGw:                   a,
		traGw:                   t,
		eventDistpatcher:        ed,
		transactionCreatedEvent: e,
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
	output := &CreateAccountOutputDTO{TransactionID: transactino.ID}

	uc.transactionCreatedEvent.SetPayload(output)
	uc.eventDistpatcher.Dispatch(uc.transactionCreatedEvent)

	return output, nil
}
