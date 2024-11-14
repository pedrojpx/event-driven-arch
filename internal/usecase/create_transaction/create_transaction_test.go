package createtransaction

import (
	"testing"

	"github.com/pedrojpx/ms-wallet/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TransactionGatewayMock struct {
	mock.Mock
}

func (m *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *AccountGatewayMock) FindByID(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func TestCreateTransactionUsecase_Execute(t *testing.T) {
	c1, _ := entity.NewClient("a", "@")
	acc1 := entity.NewAccount(c1)
	acc1.Credit(1000)
	c2, _ := entity.NewClient("b", "@")
	acc2 := entity.NewAccount(c2)
	acc2.Credit(1000)

	mockAccountGW := &AccountGatewayMock{}
	mockAccountGW.On("FindByID", acc1.ID).Return(acc1, nil)
	mockAccountGW.On("FindByID", acc2.ID).Return(acc2, nil)

	mockTransactionGW := &TransactionGatewayMock{}
	mockTransactionGW.On("Create", mock.Anything).Return(nil)

	input := CreateTrasactionInputDTO{
		AccountIDFrom: acc1.ID,
		AccountIDTo:   acc2.ID,
		Amount:        500.0,
	}

	uc := NewCreateTransactionUseCase(mockAccountGW, mockTransactionGW)
	output, err := uc.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	mockAccountGW.AssertExpectations(t)
	mockAccountGW.AssertNumberOfCalls(t, "FindByID", 2)
	mockTransactionGW.AssertExpectations(t)
	mockTransactionGW.AssertNumberOfCalls(t, "Create", 1)
}
