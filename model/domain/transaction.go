package domain

import "time"

type Transaction struct {
	Id              int
	Email           string
	Amount          int
	TransactionType string
	CreatedOn       time.Time
}
