package app

import (
	"ContractSIMSPPOB/helper"
	"database/sql"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:terserah123@localhost:5432/db_contractsimsppob?sslmode=disable")
	helper.PanicIFError(err)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db

}
