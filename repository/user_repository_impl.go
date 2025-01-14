package repository

import (
	"ContractSIMSPPOB/helper"
	"ContractSIMSPPOB/model/domain"
	"context"
	"database/sql"
	"errors"
)

type UserRepositoryImpl struct{}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (u *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "INSERT INTO users(email, firs_name, last_name, password, profile_image, saldo) VALUES($1,$2,$3,$4,$5,$6) RETURNING id"
	var id int
	err := tx.QueryRowContext(ctx, SQL, user.Email, user.FirsName, user.LastName, user.Password, user.ProfileImage, user.Saldo).Scan(&id)
	helper.PanicIFError(err)
	user.Id = id
	return user
}
func (u *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	SQL := "SELECT id, email,firs_name, last_name, password, profile_image FROM users"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIFError(err)
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.Id, &user.Email, &user.FirsName, &user.LastName, &user.Password, &user.ProfileImage)
		helper.PanicIFError(err)
		users = append(users, user)
	}
	return users
}
func (u *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
	SQL := "SELECT id, email,firs_name, last_name, password, profile_image FROM users WHERE id = $1"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIFError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Email, &user.FirsName, &user.LastName, &user.Password, &user.ProfileImage)
		helper.PanicIFError(err)
		return user, nil
	} else {
		return user, errors.New("user id not found")
	}
}
func (u *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "UPDATE users SET firs_name = $1, last_name = $2 WHERE id = $3"
	_, err := tx.ExecContext(ctx, SQL, user.FirsName, user.LastName, user.Id)
	helper.PanicIFError(err)
	return user
}
func (u *UserRepositoryImpl) UpdateImage(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "UPDATE users SET profile_image= $1 WHERE id = $2"
	_, err := tx.ExecContext(ctx, SQL, user.ProfileImage, user.Id)
	helper.PanicIFError(err)
	return user
}
func (u *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.User) {
	SQL := "DELETE FROM users WHERE id = $1"
	_, err := tx.ExecContext(ctx, SQL, user.Id)
	helper.PanicIFError(err)
}

func (u *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error) {
	SQL := "SELECT id, email,firs_name, last_name, password, profile_image FROM users WHERE email = $1"
	rows, err := tx.QueryContext(ctx, SQL, email)
	helper.PanicIFError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Email, &user.FirsName, &user.LastName, &user.Password, &user.ProfileImage)
		helper.PanicIFError(err)
		return user, nil
	} else {
		return user, errors.New("user id not found")
	}
}

func (u *UserRepositoryImpl) BalanceByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error) {
	SQL := "SELECT saldo FROM users WHERE email = $1"
	row, err := tx.QueryContext(ctx, SQL, email)
	helper.PanicIFError(err)
	defer row.Close()

	user := domain.User{}
	if row.Next() {
		err := row.Scan(&user.Saldo)
		helper.PanicIFError(err)
		return user, nil
	} else {
		return user, errors.New("email not found")
	}
}
func (u *UserRepositoryImpl) Topup(ctx context.Context, tx *sql.Tx, email string, amount int) error {
	SQL := "UPDATE users SET saldo = saldo + $1 WHERE email = $2"
	_, err := tx.ExecContext(ctx, SQL, amount, email)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepositoryImpl) SaveTransaction(ctx context.Context, tx *sql.Tx, email string, amount int, transactionType string, createdOn string) error {
	SQL := "INSERT INTO transactions (email, amount, transaction_type, created_on) VALUES ($1, $2, $3, $4)"
	_, err := tx.ExecContext(ctx, SQL, email, amount, transactionType, createdOn)
	if err != nil {
		return err
	}
	return nil
}
