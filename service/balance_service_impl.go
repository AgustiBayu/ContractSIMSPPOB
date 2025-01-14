package service

import (
	"ContractSIMSPPOB/exception"
	"ContractSIMSPPOB/helper"
	"ContractSIMSPPOB/model/web"
	"ContractSIMSPPOB/repository"
	"context"
	"database/sql"
	"errors"
	"time"

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

func (service *BalanceServiceImpl) TopUpSaldo(ctx context.Context, email string, request web.TopupCreateRequest) web.BalanceResponse {
	tx, err := service.DB.Begin()
	helper.PanicIFError(err)
	defer helper.RollbackOrCommit(tx)
	if request.TopupAmoun <= 0 {
		panic(errors.New("amount must be greater than 0"))
	}
	err = service.UserRepository.Topup(ctx, tx, email, request.TopupAmoun)
	helper.PanicIFError(err)
	err = service.UserRepository.SaveTransaction(ctx, tx, email, request.TopupAmoun, "TOPUP", time.Now().Format("2006-01-02 15:04:05"))
	helper.PanicIFError(err)
	balance, err := service.UserRepository.BalanceByEmail(ctx, tx, email)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToBalanceResponse(balance)
}
