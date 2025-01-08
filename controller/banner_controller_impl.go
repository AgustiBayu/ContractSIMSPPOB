package controller

import (
	"ContractSIMSPPOB/helper"
	"ContractSIMSPPOB/model/web"
	"ContractSIMSPPOB/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type BannerControllerImpl struct {
	BannerService service.BannerService
}

func NewBannerController(bannerService service.BannerService) BannerController {
	return &BannerControllerImpl{
		BannerService: bannerService,
	}
}
func (controller *BannerControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bannerResponse := controller.BannerService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:    http.StatusOK,
		Message: "sukses",
		Data:    bannerResponse,
	}
	helper.WriteResponseBody(writer, webResponse)
}
