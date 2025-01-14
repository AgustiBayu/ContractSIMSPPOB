package repository

import (
	"ContractSIMSPPOB/helper"
	"ContractSIMSPPOB/model/domain"
	"context"
	"database/sql"
	"errors"
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

func (l *LayananRepositoryImpl) GetByServiceCode(ctx context.Context, tx *sql.Tx, serviceCode string) (domain.Layanan, error) {
	SQL := "SELECT id, service_code, service_name, service_icon, service_tarif FROM layanans WHERE service_code = $1"
	row, err := tx.QueryContext(ctx, SQL, serviceCode)
	helper.PanicIFError(err)
	defer row.Close()

	layanan := domain.Layanan{}
	if row.Next() {
		err := row.Scan(&layanan.Id, &layanan.ServiceCode, &layanan.ServiceIcon, &layanan.ServiceName, &layanan.ServiceTarif)
		helper.PanicIFError(err)
		return layanan, nil
	} else {
		return layanan, errors.New("service code not found")
	}
}
