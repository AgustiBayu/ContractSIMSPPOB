package controller

import (
	"ContractSIMSPPOB/helper"
	"ContractSIMSPPOB/model/web"
	"ContractSIMSPPOB/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userRegisterRequest := web.UserRegisterRequst{}
	helper.ReadRequstBody(request, &userRegisterRequest)

	userResponse := controller.UserService.Register(request.Context(), userRegisterRequest)
	webResponse := web.WebResponse{
		Code:    http.StatusCreated,
		Message: "registrasi berhasil",
		Data:    userResponse,
	}
	helper.WriteResponseBody(writer, webResponse)
}
func (controller *UserControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userLoginRequest := web.UserLoginRequst{}
	helper.ReadRequstBody(request, &userLoginRequest)

	userResponse, err := controller.UserService.Login(request.Context(), userLoginRequest)
	if err != nil {
		helper.PanicIFError(err)
	}
	webResponse := web.WebResponse{
		Code:    http.StatusCreated,
		Message: "registrasi berhasil",
		Data:    userResponse,
	}
	helper.WriteResponseBody(writer, webResponse)
}
