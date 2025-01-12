package controller

import (
	"ContractSIMSPPOB/helper"
	"ContractSIMSPPOB/model/web"
	"ContractSIMSPPOB/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type LayananControllerImpl struct {
	LayananService service.LayananService
}

func NewLayananController(layananService service.LayananService) LayananController {
	return &LayananControllerImpl{
		LayananService: layananService,
	}
}
func (controller *LayananControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	layananResponse := controller.LayananService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:    http.StatusOK,
		Message: "sukses",
		Data:    layananResponse,
	}
	helper.WriteResponseBody(writer, webResponse)
}
