package repository

import (
	"ContractSIMSPPOB/helper"
	"ContractSIMSPPOB/model/domain"
	"context"
	"database/sql"
)

type BannerRepositoryImpl struct{}

func NewBannerRepository() BannerRepository {
	return &BannerRepositoryImpl{}
}

func (b *BannerRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Banner {
	SQL := "SELECT id, banner_name, banner_image, description FROM banners"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIFError(err)
	defer rows.Close()

	var banners []domain.Banner
	for rows.Next() {
		banner := domain.Banner{}
		err := rows.Scan(&banner.Id, &banner.BannerName, &banner.BannerImage, &banner.Description)
		helper.PanicIFError(err)
		banners = append(banners, banner)
	}
	return banners
}
