package repository

import (
	"ContractSIMSPPOB/helper"
	"ContractSIMSPPOB/model/domain"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type TransactionRespositoryImpl struct{}

func NewTransactionRepository() TransactionRespository {
	return &TransactionRespositoryImpl{}
}

func (t *TransactionRespositoryImpl) FindByServiceCode(ctx context.Context, tx *sql.Tx, serviceCode string) (domain.Layanan, domain.Transaction, error) {
	SQL := "select l.service_code, service_name, t.transaction_type, t.created_on from transactions t inner join layanans l on t.service_code_fk = l.service_code WHERE t.service_code_fk = $1"
	row, err := tx.QueryContext(ctx, SQL, serviceCode)
	helper.PanicIFError(err)
	defer row.Close()

	layanan := domain.Layanan{}
	transaction := domain.Transaction{}
	if row.Next() {
		err := row.Scan(&layanan.ServiceCode, &layanan.ServiceName, &transaction.TransactionType, &transaction.CreatedOn)
		helper.PanicIFError(err)
		return layanan, transaction, nil
	} else {
		return layanan, transaction, errors.New("service code not found")
	}
}
func (t *TransactionRespositoryImpl) CreateTransaction(ctx context.Context, tx *sql.Tx, transaction domain.Transaction) error {
	SQL := "INSERT INTO transactions (email, amount, transaction_type, created_on) VALUES ($1, $2, $3, $4)"
	_, err := tx.ExecContext(ctx, SQL, transaction.Email, transaction.Amount, transaction.TransactionType, transaction.CreatedOn)
	if err != nil {
		return err
	}
	return nil
}
func (t *TransactionRespositoryImpl) CheckBalance(ctx context.Context, tx *sql.Tx, email string) (int, error) {
	SQL := "SELECT saldo FROM users WHERE email = $1"
	var balance int

	// Use QueryRowContext to get a single result
	err := tx.QueryRowContext(ctx, SQL, email).Scan(&balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("no user found with email: %s", email)
		}
		return 0, fmt.Errorf("failed to retrieve balance: %w", err)
	}

	return balance, nil
}
