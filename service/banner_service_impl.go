package service

import (
	"ContractSIMSPPOB/helper"
	"ContractSIMSPPOB/model/web"
	"ContractSIMSPPOB/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type BannerServiceImpl struct {
	BannerRepository repository.BannerRepository
	DB               *sql.DB
	Validate         *validator.Validate
}

func NewBannerService(bannerRepository repository.BannerRepository,
	DB *sql.DB, validate *validator.Validate) BannerService {
	return &BannerServiceImpl{
		BannerRepository: bannerRepository,
		DB:               DB,
		Validate:         validate,
	}
}

func (service *BannerServiceImpl) FindAll(ctx context.Context) []web.BannerResponse {
	tx, err := service.DB.Begin()
	helper.PanicIFError(err)
	defer helper.RollbackOrCommit(tx)
	banner := service.BannerRepository.FindAll(ctx, tx)
	return helper.ToBannerResponses(banner)
}
