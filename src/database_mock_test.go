package src_test

import "github.com/stretchr/testify/mock"
import . "github.com/dating-api/src"

type databaseMock struct {
	Mock *mock.Mock
}

func (d *databaseMock) GetUserByEmail(email string) (*User, error) {
	args := d.Mock.Called(email)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*User), args.Error(1)
}

func (d *databaseMock) StoreUser(user User) (*int64, error) {
	args := d.Mock.Called(user)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*int64), args.Error(1)
}
