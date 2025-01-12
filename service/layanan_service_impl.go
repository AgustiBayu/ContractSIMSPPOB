package service

import (
	"ContractSIMSPPOB/helper"
	"ContractSIMSPPOB/model/web"
	"ContractSIMSPPOB/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type LayananServiceImpl struct {
	LayananRepository repository.LayananRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewLayananService(layananRepository repository.LayananRepository,
	DB *sql.DB, validate *validator.Validate) LayananService {
	return &LayananServiceImpl{
		LayananRepository: layananRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service *LayananServiceImpl) FindAll(ctx context.Context) []web.LayananResponse {
	tx, err := service.DB.Begin()
	helper.PanicIFError(err)
	defer helper.RollbackOrCommit(tx)
	layanan := service.LayananRepository.FindAll(ctx, tx)
	return helper.ToLayananResponses(layanan)
}
