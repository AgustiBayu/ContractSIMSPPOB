package service

import (
	"ContractSIMSPPOB/model/web"
	"context"
)

type LayananService interface {
	FindAll(ctx context.Context) []web.LayananResponse
}
