package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserProfileController interface {
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateImage(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
