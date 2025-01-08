package repository

import (
	"ContractSIMSPPOB/model/domain"
	"context"
	"database/sql"
)

type BannerRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Banner
}
