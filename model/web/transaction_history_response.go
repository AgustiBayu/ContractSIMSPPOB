package web

type TransactionHistory struct {
	Id              int    `json:"id"`
	Email           string `json:"email"`
	Amount          int    `json:"amount"`
	TransactionType string `json:"transaction_type"`
	CreatedOn       string `json:"created_on"`
}
