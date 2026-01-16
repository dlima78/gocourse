package mock_user_controller

import (
	rest_err "github.com/dlima78/gocourse/src/configuration"
	"github.com/dlima78/gocourse/src/model"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUserService(u model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	args := m.Called(u)
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
func (m *MockUserService) UpdateUserService(id string, u model.UserDomainInterface) *rest_err.RestErr {
	args := m.Called(id, u)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*rest_err.RestErr)
}
func (m *MockUserService) FindUserByEmailService(email string) (model.UserDomainInterface, *rest_err.RestErr) {
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
func (m *MockUserService) FindUserByIDService(id string) (model.UserDomainInterface, *rest_err.RestErr) {
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
func (m *MockUserService) DeleteUserService(id string) *rest_err.RestErr {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*rest_err.RestErr)
}
func (m *MockUserService) LoginUserService(u model.UserDomainInterface) (model.UserDomainInterface, string, *rest_err.RestErr) {
	args := m.Called(u)
	if args.Get(0) == nil {
		if args.Get(2) == nil {
			return nil, "", nil
		}
		return nil, "", args.Get(2).(*rest_err.RestErr)
	}
	if args.Get(2) == nil {
		return args.Get(0).(model.UserDomainInterface), args.String(1), nil
	}
	return args.Get(0).(model.UserDomainInterface), args.String(1), args.Get(2).(*rest_err.RestErr)
}
