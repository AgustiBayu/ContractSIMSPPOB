package service

import (
	"ContractSIMSPPOB/model/web"
	"context"
)

type TransactionService interface {
	ProcessTransaction(ctx context.Context, request web.TransactionCreateRequest, email string) web.TransactionResponse
}
