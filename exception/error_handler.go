package exception

import (
	"ContractSIMSPPOB/helper"
	"ContractSIMSPPOB/model/web"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if validationError(writer, request, err) {
		return
	}
	if notFoundError(writer, request, err) {
		return
	}
	internalServerError(writer, request, err)
}

func HandleNotFound(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusNotFound)
	webResponse := web.ErrorResponse{
		Code:    http.StatusNotFound,
		Message: "end point is not valid",
	}
	helper.WriteResponseBody(writer, webResponse)
}

func validationError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := web.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: exception.Error(),
		}
		helper.WriteResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: exception.Error,
		}
		helper.WriteResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}
func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	var errorMessage string
	if err == nil {
		errorMessage = "An unexpected error occurred"
	} else {
		switch e := err.(type) {
		case string:
			errorMessage = e
		case error:
			errorMessage = e.Error()
		default:
			errorMessage = "An unexpected error occurred"
		}
	}

	webResponse := web.ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: errorMessage,
	}
	helper.WriteResponseBody(writer, webResponse)
}
