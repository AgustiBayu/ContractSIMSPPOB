package controller

import (
	"ContractSIMSPPOB/helper"
	"ContractSIMSPPOB/model/web"
	"ContractSIMSPPOB/service"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/julienschmidt/httprouter"
)

type BalanceControllerImpl struct {
	BalanceService service.BalanceService
}

func NewBalanceController(balanceService service.BalanceService) BalanceController {
	return &BalanceControllerImpl{
		BalanceService: balanceService,
	}
}

func (controller *BalanceControllerImpl) GetBalanceByEmail(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	tokenString := request.Header.Get("Authorization")[7:] // Remove "Bearer " prefix
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(helper.GetSecretKey()), nil
	})
	helper.PanicIFError(err)
	claims := token.Claims.(jwt.MapClaims)
	email := claims["email"].(string)
	balanceResponse := controller.BalanceService.GetBalanceByEmail(request.Context(), email)
	webResponse := web.WebResponse{
		Code:    http.StatusOK,
		Message: "Get Balance Berhasil",
		Data:    balanceResponse,
	}
	helper.WriteResponseBody(writer, webResponse)
}
