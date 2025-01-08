package service

import (
	"ContractSIMSPPOB/exception"
	"ContractSIMSPPOB/helper"
	"ContractSIMSPPOB/model/web"
	"ContractSIMSPPOB/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type UserProfileServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserProfileService(userRepository repository.UserRepository, DB *sql.DB,
	validate *validator.Validate) UserProfileService {
	return &UserProfileServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *UserProfileServiceImpl) FindAll(ctx context.Context) []web.UserProfileResponse {
	tx, err := service.DB.Begin()
	helper.PanicIFError(err)
	defer helper.RollbackOrCommit(tx)
	user := service.UserRepository.FindAll(ctx, tx)
	return helper.ToUserProfileResponses(user)
}
func (service *UserProfileServiceImpl) Update(ctx context.Context, request web.UserProfileUpdateRequest) web.UserProfileResponse {
	tx, err := service.DB.Begin()
	helper.PanicIFError(err)
	defer helper.RollbackOrCommit(tx)
	user, err := service.UserRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	user.FirsName = request.FirsName
	user.LastName = request.LastName
	user = service.UserRepository.Update(ctx, tx, user)
	return helper.ToUserProfileResponse(user)
}
func (service *UserProfileServiceImpl) UpdateImage(ctx context.Context, request web.UserProfileUpdateImageRequest) web.UserProfileResponse {
	tx, err := service.DB.Begin()
	helper.PanicIFError(err)
	defer helper.RollbackOrCommit(tx)

	user, err := service.UserRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	user.ProfileImage = request.ProfileImage
	user = service.UserRepository.UpdateImage(ctx, tx, user)
	return helper.ToUserProfileResponse(user)
}
