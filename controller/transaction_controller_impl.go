package controller

import (
	"ContractSIMSPPOB/helper"
	"ContractSIMSPPOB/model/web"
	"ContractSIMSPPOB/service"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/julienschmidt/httprouter"
)

type TransactionControllerImpl struct {
	TransactionService service.TransactionService
}

func NewTransactionController(transactionService service.TransactionService) TransactionController {
	return &TransactionControllerImpl{
		TransactionService: transactionService,
	}
}

func (controller *TransactionControllerImpl) ProcessTransaction(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transactionCreateRequest := web.TransactionCreateRequest{}
	helper.ReadRequstBody(request, &transactionCreateRequest)
	tokenString := request.Header.Get("Authorization")[7:] // Remove "Bearer " prefix
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(helper.GetSecretKey()), nil
	})
	helper.PanicIFError(err)
	claims := token.Claims.(jwt.MapClaims)
	email := claims["email"].(string)
	transactionResponse := controller.TransactionService.ProcessTransaction(request.Context(), transactionCreateRequest, email)
	webResponse := web.WebResponse{
		Code:    http.StatusOK,
		Message: "transaksi berhasil",
		Data:    transactionResponse,
	}
	helper.WriteResponseBody(writer, webResponse)
}
