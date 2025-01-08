package main

import (
	"ContractSIMSPPOB/app"
	"ContractSIMSPPOB/controller"
	"ContractSIMSPPOB/exception"
	"ContractSIMSPPOB/helper"
	"ContractSIMSPPOB/repository"
	"ContractSIMSPPOB/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

func main() {
	validate := validator.New()
	db := app.NewDB()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	router := httprouter.New()
	router.POST("/api/register", userController.Register)
	router.POST("/api/login", userController.Login)

	router.NotFound = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		exception.HandleNotFound(writer, request)
	})
	router.PanicHandler = func(writer http.ResponseWriter, request *http.Request, err interface{}) {
		exception.ErrorHandler(writer, request, err)
	}

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}
	err := server.ListenAndServe()
	helper.PanicIFError(err)
}
