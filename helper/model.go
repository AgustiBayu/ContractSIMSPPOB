package helper

import (
	"ContractSIMSPPOB/model/domain"
	"ContractSIMSPPOB/model/web"
)

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
