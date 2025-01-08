package service

import (
	"ContractSIMSPPOB/model/web"
	"context"
)

type UserProfileService interface {
	FindAll(ctx context.Context) []web.UserProfileResponse
	Update(ctx context.Context, request web.UserProfileUpdateRequest) web.UserProfileResponse
	UpdateImage(ctx context.Context, request web.UserProfileUpdateImageRequest) web.UserProfileResponse
}
