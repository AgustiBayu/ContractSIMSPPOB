package controller

import (
	"ContractSIMSPPOB/helper"
	"ContractSIMSPPOB/model/web"
	"ContractSIMSPPOB/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type UserProfileControllerImpl struct {
	UserProfileService service.UserProfileService
}

func NewUserProfileController(userProfileService service.UserProfileService) UserProfileController {
	return &UserProfileControllerImpl{
		UserProfileService: userProfileService,
	}
}

func (controller *UserProfileControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userResponse := controller.UserProfileService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:    http.StatusOK,
		Message: "sukses",
		Data:    userResponse,
	}
	helper.WriteResponseBody(writer, webResponse)
}
func (controller *UserProfileControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userProfileUpdateRequest := web.UserProfileUpdateRequest{}
	helper.ReadRequstBody(request, &userProfileUpdateRequest)
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIFError(err)

	userProfileUpdateRequest.Id = id
	userResponse := controller.UserProfileService.Update(request.Context(), userProfileUpdateRequest)
	webResponse := web.WebResponse{
		Code:    http.StatusOK,
		Message: "update profile berhasil",
		Data:    userResponse,
	}
	helper.WriteResponseBody(writer, webResponse)
}
func (controller *UserProfileControllerImpl) UpdateImage(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userProfileUpdateImageRequest := web.UserProfileUpdateImageRequest{}
	helper.ReadRequstBody(request, &userProfileUpdateImageRequest)
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIFError(err)

	userProfileUpdateImageRequest.Id = id
	userResponse := controller.UserProfileService.UpdateImage(request.Context(), userProfileUpdateImageRequest)
	webResponse := web.WebResponse{
		Code:    http.StatusOK,
		Message: "update profile image berhasil",
		Data:    userResponse,
	}
	helper.WriteResponseBody(writer, webResponse)
}
