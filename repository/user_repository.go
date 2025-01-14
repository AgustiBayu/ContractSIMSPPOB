package repository

import (
	"ContractSIMSPPOB/model/domain"
	"context"
	"database/sql"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	FindAll(ctx context.Context, tx *sql.Tx) []domain.User
	FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error)
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error)
	Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	UpdateImage(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	BalanceByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error)
	Topup(ctx context.Context, tx *sql.Tx, email string, amount int) error
	SaveTransaction(ctx context.Context, tx *sql.Tx, email string, amount int, transactionType string, created_on string) error
	Delete(ctx context.Context, tx *sql.Tx, user domain.User)
}
