package helper

import (
	"ContractSIMSPPOB/model/domain"
	"ContractSIMSPPOB/model/web"
)

func ToBannerResponses(banners []domain.Banner) []web.BannerResponse {
	var bannerResponse []web.BannerResponse
	for _, banner := range banners {
		bannerResponse = append(bannerResponse, ToBannerResponse(banner))
	}
	return bannerResponse
}

func ToBannerResponse(banner domain.Banner) web.BannerResponse {
	return web.BannerResponse{
		Id:          banner.Id,
		BannerName:  banner.BannerName,
		BannerImage: banner.BannerImage,
		Description: banner.Description,
	}
}

func ToUserProfileResponses(users []domain.User) []web.UserProfileResponse {
	var userResponse []web.UserProfileResponse
	for _, user := range users {
		userResponse = append(userResponse, ToUserProfileResponse(user))
	}
	return userResponse
}

func ToUserProfileResponse(user domain.User) web.UserProfileResponse {
	return web.UserProfileResponse{
		Id:           user.Id,
		Email:        user.Email,
		FirsName:     user.FirsName,
		LastName:     user.LastName,
		ProfileImage: user.ProfileImage,
	}
}
