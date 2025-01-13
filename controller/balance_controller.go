package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type BalanceController interface {
	GetBalanceByEmail(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	TopUpSaldo(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
