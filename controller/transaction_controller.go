package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type TransactionController interface {
	ProcessTransaction(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
