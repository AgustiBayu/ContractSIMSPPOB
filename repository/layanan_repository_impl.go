package repository

import (
	"ContractSIMSPPOB/helper"
	"ContractSIMSPPOB/model/domain"
	"context"
	"database/sql"
)

type LayananRepositoryImpl struct{}

func NewLayananRepository() LayananRepository {
	return &LayananRepositoryImpl{}
}

func (l *LayananRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Layanan {
	SQL := "SELECT id, service_code, service_name, service_icon, service_tarif FROM layanans"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIFError(err)
	defer rows.Close()

	var layanans []domain.Layanan
	for rows.Next() {
		layanan := domain.Layanan{}
		err := rows.Scan(&layanan.Id, &layanan.ServiceCode, &layanan.ServiceName, &layanan.ServiceIcon, &layanan.ServiceTarif)
		helper.PanicIFError(err)
		layanans = append(layanans, layanan)
	}
	return layanans
}
