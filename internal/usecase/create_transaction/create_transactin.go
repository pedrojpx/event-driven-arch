package createtransaction

import (
	"context"

	"github.com/pedrojpx/ms-wallet/internal/entity"
	"github.com/pedrojpx/ms-wallet/internal/gateway"
	"github.com/pedrojpx/ms-wallet/pkg/events"
	"github.com/pedrojpx/ms-wallet/pkg/uow"
)

type CreateTrasactionInputDTO struct {
	AccountIDFrom string  `json:"account_from"`
	AccountIDTo   string  `json:"account_to"`
	Amount        float64 `json:"amount"`
}

type CreateTransactionOutputDTO struct {
	TransactionID string  `json:"transaction_id"`
	AccountIDFrom string  `json:"account_from"`
	AccountIDTo   string  `json:"account_to"`
	Amount        float64 `json:"amount"`
}

type CreateTransactionUsecase struct {
	uow                     uow.UowInterface
	eventDistpatcher        events.EventDispatcherInterface
	transactionCreatedEvent events.EventInterface
}

func NewCreateTransactionUseCase(u uow.UowInterface, ed events.EventDispatcherInterface, e events.EventInterface) *CreateTransactionUsecase {
	return &CreateTransactionUsecase{
		uow:                     u,
		eventDistpatcher:        ed,
		transactionCreatedEvent: e,
	}
}

func (uc *CreateTransactionUsecase) Execute(ctx context.Context, input CreateTrasactionInputDTO) (*CreateTransactionOutputDTO, error) {
	output := &CreateTransactionOutputDTO{}
	err := uc.uow.Do(ctx, func(_ *uow.Uow) error {
		accRepo := uc.getAccountRepo(ctx)
		transRepo := uc.getTransactionRepo(ctx)

		from, err := accRepo.FindByID(input.AccountIDFrom)
		if err != nil {
			return err
		}
		to, err := accRepo.FindByID(input.AccountIDTo)
		if err != nil {
			return err
		}
		transactino, err := entity.NewTransaction(from, to, input.Amount)
		if err != nil {
			return err
		}
		err = transRepo.Create(transactino)
		if err != nil {
			return err
		}
		err = accRepo.UpdateBalance(from)
		err = accRepo.UpdateBalance(to)
		output.TransactionID = transactino.ID
		output.AccountIDFrom = transactino.From.Client.ID
		output.AccountIDTo = transactino.To.Client.ID
		output.Amount = transactino.Amount
		return nil
	})

	if err != nil {
		return nil, err
	}

	uc.transactionCreatedEvent.SetPayload(output)
	uc.eventDistpatcher.Dispatch(uc.transactionCreatedEvent)

	return output, nil
}

func (uc *CreateTransactionUsecase) getAccountRepo(ctx context.Context) gateway.AccountGateway {
	repo, err := uc.uow.GetRepository(ctx, "AccountDB")
	if err != nil {
		panic(err)
	}
	return repo.(gateway.AccountGateway)
}

func (uc *CreateTransactionUsecase) getTransactionRepo(ctx context.Context) gateway.TransactionGateway {
	repo, err := uc.uow.GetRepository(ctx, "TransactionDB")
	if err != nil {
		panic(err)
	}
	return repo.(gateway.TransactionGateway)
}
