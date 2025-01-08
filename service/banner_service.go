package service

import (
	"ContractSIMSPPOB/model/web"
	"context"
)

type BannerService interface {
	FindAll(ctx context.Context) []web.BannerResponse
}
