package repository

import (
	"ContractSIMSPPOB/model/domain"
	"context"
	"database/sql"
)

type LayananRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Layanan
	GetByServiceCode(ctx context.Context, tx *sql.Tx, serviceCode string) (domain.Layanan, error)
}
