package service

import (
	"ContractSIMSPPOB/model/web"
	"context"
)

type UserService interface {
	Register(ctx context.Context, request web.UserRegisterRequst) error
	Login(ctx context.Context, request web.UserLoginRequst) (web.UserResponse, error)
}
