package web

type TransactionResponse struct {
	InvoiceNumber   string `json:"invoice_number"`
	ServiceCode     string `json:"service_code"`
	ServiceName     string `json:"service_name"`
	Amount          int    `json:"total_amount"`
	TransactionType string `json:"transaction_type"`
	CreatedOn       string `json:"create_on"`
}
