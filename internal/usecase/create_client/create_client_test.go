package createclient

import (
	"testing"

	"github.com/pedrojpx/ms-wallet/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ClientGatewayMock struct {
	mock.Mock
}

func (m *ClientGatewayMock) Save(client *entity.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

func (m *ClientGatewayMock) FindByID(id string) (*entity.Client, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
}

func TestCreateClientUseCase_Execute(t *testing.T) {
	m := &ClientGatewayMock{}
	m.On("FindByID", mock.Anything).Return(nil)

	uc := NewCreateClientUseCase(m)
	output, err := uc.Execute(CreateClientInputDTO{Name: "a", Email: "@"})

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	assert.Equal(t, output.Name, "a")
	assert.Equal(t, output.Email, "@")
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Save", 1)
}
