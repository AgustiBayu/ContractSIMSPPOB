package helper

import (
	"ContractSIMSPPOB/model/domain"
	"ContractSIMSPPOB/model/web"
	"ContractSIMSPPOB/utils"
)

func ToLayananResponses(banners []domain.Layanan) []web.LayananResponse {
	var layananResponse []web.LayananResponse
	for _, banner := range banners {
		layananResponse = append(layananResponse, web.LayananResponse(banner))
	}
	return layananResponse
}

func ToLayananResponse(banner domain.Layanan) web.LayananResponse {
	return web.LayananResponse{
		Id:           banner.Id,
		ServiceCode:  banner.ServiceCode,
		ServiceName:  banner.ServiceName,
		ServiceIcon:  banner.ServiceIcon,
		ServiceTarif: banner.ServiceTarif,
	}
}
func ToTransactionResponse(layanan domain.Layanan, transaction domain.Transaction) web.TransactionResponse {
	return web.TransactionResponse{
		InvoiceNumber:   utils.GenerateInvoiceNumber(),
		ServiceCode:     layanan.ServiceCode,
		ServiceName:     layanan.ServiceName,
		Amount:          transaction.Amount,
		TransactionType: transaction.TransactionType,
		CreatedOn:       FormatTanggal(transaction.CreatedOn),
	}
}
func ToBalanceResponse(saldo domain.User) web.BalanceResponse {
	return web.BalanceResponse{
		Saldo: saldo.Saldo,
	}
}

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
