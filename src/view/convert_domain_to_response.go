package view

import (
	"github.com/dlima78/gocourse/src/controller/model/response"
	"github.com/dlima78/gocourse/src/model"
)

func ConvertDomainToResponse(userDomain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:    "",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	} // to be implement
}
