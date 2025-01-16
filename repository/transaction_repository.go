package repository

import (
	"ContractSIMSPPOB/model/domain"
	"context"
	"database/sql"
)

type TransactionRespository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Transaction
	FindByServiceCode(ctx context.Context, tx *sql.Tx, serviceCode string) (domain.Layanan, domain.Transaction, error)
	CreateTransaction(ctx context.Context, tx *sql.Tx, transaction domain.Transaction) error
	CheckBalance(ctx context.Context, tx *sql.Tx, email string) (int, error)
}
