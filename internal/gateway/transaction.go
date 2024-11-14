package gateway

import "github.com/pedrojpx/ms-wallet/internal/entity"

type TransactionGateway interface {
	Create(transactino *entity.Transaction) error
}
