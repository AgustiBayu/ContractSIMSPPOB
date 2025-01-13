package service

import (
	"ContractSIMSPPOB/exception"
	"ContractSIMSPPOB/helper"
	"ContractSIMSPPOB/model/web"
	"ContractSIMSPPOB/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type BalanceServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewBalanceService(userRepository repository.UserRepository,
	DB *sql.DB, validate *validator.Validate) BalanceService {
	return &BalanceServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}
func (service *BalanceServiceImpl) GetBalanceByEmail(ctx context.Context, email string) web.BalanceResponse {
	tx, err := service.DB.Begin()
	helper.PanicIFError(err)
	defer helper.RollbackOrCommit(tx)
	balance, err := service.UserRepository.BalanceByEmail(ctx, tx, email)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToBalanceResponse(balance)
}
