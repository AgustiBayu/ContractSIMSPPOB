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
	SQL := "INSERT INTO users(email, firs_name, last_name, password) VALUES($1,$2,$3,$4) RETURNING id"
	var id int
	err := tx.QueryRowContext(ctx, SQL, user.Email, user.FirsName, user.LastName, user.Password).Scan(&id)
	helper.PanicIFError(err)
	user.Id = id
	return user
}
func (u *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	SQL := "SELECT id, email,firs_name, last_name, password FROM users"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIFError(err)
	defer rows.Close()

	var users []domain.User
	if rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.Id, &user.Email, &user.FirsName, &user.LastName, &user.Password)
		helper.PanicIFError(err)
		users = append(users, user)
	}
	return users
}
func (u *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
	SQL := "SELECT id, email,firs_name, last_name, password FROM users WHERE id = $1"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIFError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Email, &user.FirsName, &user.LastName, &user.Password)
		helper.PanicIFError(err)
		return user, nil
	} else {
		return user, errors.New("user id not found")
	}
}
func (u *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "UPDATE users SET email = $1,firs_name = $2, last_name = $3, password = $4 WHERE id = $5"
	_, err := tx.ExecContext(ctx, SQL, user.Email, user.FirsName, user.LastName, user.Password)
	helper.PanicIFError(err)
	return user
}
func (u *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.User) {
	SQL := "DELETE FROM users WHERE id = $1"
	_, err := tx.ExecContext(ctx, SQL, user.Id)
	helper.PanicIFError(err)
}

func (u *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error) {
	SQL := "SELECT id, email,firs_name, last_name, password FROM users WHERE email = $1"
	rows, err := tx.QueryContext(ctx, SQL, email)
	helper.PanicIFError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Email, &user.FirsName, &user.LastName, &user.Password)
		helper.PanicIFError(err)
		return user, nil
	} else {
		return user, errors.New("user id not found")
	}
}
