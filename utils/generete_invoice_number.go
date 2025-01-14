package utils

import "time"

func GenerateInvoiceNumber() string {
	return "INV" + time.Now().Format("20060102-150405")
}
