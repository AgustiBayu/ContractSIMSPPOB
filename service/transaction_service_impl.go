package service

import (
	"ContractSIMSPPOB/exception"
	"ContractSIMSPPOB/helper"
	"ContractSIMSPPOB/model/domain"
	"ContractSIMSPPOB/model/web"
	"ContractSIMSPPOB/repository"
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
)

type TransactionServiceImpl struct {
	TransactionRepository repository.TransactionRespository
	LayananRepository     repository.LayananRepository
	DB                    *sql.DB
	Validate              *validator.Validate
}

func NewTransactionService(transactionRepository repository.TransactionRespository, layananRepository repository.LayananRepository,
	DB *sql.DB, validate *validator.Validate) TransactionService {
	return &TransactionServiceImpl{
		TransactionRepository: transactionRepository,
		LayananRepository:     layananRepository,
		DB:                    DB,
		Validate:              validate,
	}
}

func (service *TransactionServiceImpl) ProcessTransaction(ctx context.Context, request web.TransactionCreateRequest, email string) web.TransactionResponse {
	// Validasi request
	err := service.Validate.Struct(request)
	helper.PanicIFError(err)

	// Memulai transaksi database
	tx, err := service.DB.Begin()
	helper.PanicIFError(err)
	defer helper.RollbackOrCommit(tx)

	// Memastikan layanan tersedia
	layanan, err := service.LayananRepository.GetByServiceCode(ctx, tx, request.ServiceCode)
	if err != nil {
		panic(exception.NewNotFoundError("Service not found"))
	}

	// Mengecek saldo pengguna
	balance, err := service.TransactionRepository.CheckBalance(ctx, tx, email)
	helper.PanicIFError(err)
	if balance < layanan.ServiceTarif {
		panic(errors.New("Insufficient balance"))
	}

	// Membuat transaksi
	transaction := domain.Transaction{
		Email:           email,
		Amount:          layanan.ServiceTarif,
		TransactionType: "PAYMENT",
		CreatedOn:       time.Now(),
	}
	err = service.TransactionRepository.CreateTransaction(ctx, tx, transaction)
	helper.PanicIFError(err)

	// Mengembalikan respons
	return helper.ToTransactionResponse(layanan, transaction)
}
