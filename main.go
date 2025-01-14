package main

import (
	"ContractSIMSPPOB/app"
	"ContractSIMSPPOB/controller"
	"ContractSIMSPPOB/exception"
	"ContractSIMSPPOB/helper"
	"ContractSIMSPPOB/middleware"
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
	userProfileService := service.NewUserProfileService(userRepository, db, validate)
	userProfileController := controller.NewUserProfileController(userProfileService)
	bannerRepository := repository.NewBannerRepository()
	bannerService := service.NewBannerService(bannerRepository, db, validate)
	bannerController := controller.NewBannerController(bannerService)
	layananRepository := repository.NewLayananRepository()
	layananService := service.NewLayananService(layananRepository, db, validate)
	layananController := controller.NewLayananController(layananService)
	balanceService := service.NewBalanceService(userRepository, db, validate)
	balanceController := controller.NewBalanceController(balanceService)
	transactionRepository := repository.NewTransactionRepository()
	transactionService := service.NewTransactionService(transactionRepository, layananRepository, db, validate)
	transactionController := controller.NewTransactionController(transactionService)

	router := httprouter.New()
	router.POST("/api/register", userController.Register)
	router.POST("/api/login", userController.Login)
	router.GET("/api/profile", middleware.JWTAuth(userService, userProfileController.FindAll))
	router.PUT("/api/profile/update/:userId", userProfileController.Update)
	router.PUT("/api/profile/image/:userId", userProfileController.UpdateImage)
	router.GET("/api/banner", bannerController.FindAll)
	router.GET("/api/services", layananController.FindAll)
	router.GET("/api/balance", middleware.JWTAuth(userService, balanceController.GetBalanceByEmail))
	router.POST("/api/topup", middleware.JWTAuth(userService, balanceController.TopUpSaldo))
	router.POST("/api/transaction", transactionController.ProcessTransaction)

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
