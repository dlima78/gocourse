package mock_user_repository

import (
	rest_err "github.com/dlima78/gocourse/src/configuration"
	"github.com/dlima78/gocourse/src/model"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(user model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	args := m.Called(user)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(*rest_err.RestErr)
	}
	if args.Get(1) == nil {
		return args.Get(0).(model.UserDomainInterface), nil
	}
	return args.Get(0).(model.UserDomainInterface), args.Get(1).(*rest_err.RestErr)
}

func (m *MockUserRepository) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	args := m.Called(email)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(*rest_err.RestErr)
	}
	if args.Get(1) == nil {
		return args.Get(0).(model.UserDomainInterface), nil
	}
	return args.Get(0).(model.UserDomainInterface), args.Get(1).(*rest_err.RestErr)
}

func (m *MockUserRepository) FindUserByID(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	args := m.Called(id)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(*rest_err.RestErr)
	}
	if args.Get(1) == nil {
		return args.Get(0).(model.UserDomainInterface), nil
	}
	return args.Get(0).(model.UserDomainInterface), args.Get(1).(*rest_err.RestErr)
}

func (m *MockUserRepository) UpdateUser(id string, user model.UserDomainInterface) *rest_err.RestErr {
	args := m.Called(id, user)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*rest_err.RestErr)
}

func (m *MockUserRepository) DeleteUser(id string) *rest_err.RestErr {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*rest_err.RestErr)
}

func (m *MockUserRepository) LoginUser(user model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	args := m.Called(user)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(*rest_err.RestErr)
	}
	if args.Get(1) == nil {
		return args.Get(0).(model.UserDomainInterface), nil
	}
	return args.Get(0).(model.UserDomainInterface), args.Get(1).(*rest_err.RestErr)
}

func (m *MockUserRepository) FindUserByEmailAndPassword(email, password string) (model.UserDomainInterface, *rest_err.RestErr) {
	args := m.Called(email, password)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(*rest_err.RestErr)
	}
	if args.Get(1) == nil {
		return args.Get(0).(model.UserDomainInterface), nil
	}
	return args.Get(0).(model.UserDomainInterface), args.Get(1).(*rest_err.RestErr)
}
