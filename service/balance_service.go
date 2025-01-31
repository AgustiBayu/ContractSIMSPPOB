package service

import (
	"ContractSIMSPPOB/model/web"
	"context"
)

type BalanceService interface {
	GetBalanceByEmail(ctx context.Context, email string) web.BalanceResponse
	TopUpSaldo(ctx context.Context, email string, request web.TopupCreateRequest) web.BalanceResponse
}
