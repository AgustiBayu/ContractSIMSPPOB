package domain

import "time"

type Transaction struct {
	Id              int
	ServiceCode     string
	Email           string
	Amount          int
	TransactionType string
	CreatedOn       time.Time
}
