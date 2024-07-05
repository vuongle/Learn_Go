package db

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type AccountMock struct {
	mock.Mock
}

func (m *AccountMock) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).(Account), args.Error(1)
}
