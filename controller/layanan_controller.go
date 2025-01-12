package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type LayananController interface {
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
