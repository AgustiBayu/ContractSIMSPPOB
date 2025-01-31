package service

import (
	"ContractSIMSPPOB/helper"
	"ContractSIMSPPOB/model/domain"
	"ContractSIMSPPOB/model/web"
	"ContractSIMSPPOB/repository"
	"context"
	"database/sql"
	"errors"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(userService repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userService,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) Register(ctx context.Context, request web.UserRegisterRequst) error {
	err := service.Validate.Struct(request)
	helper.PanicIFError(err)
	tx, err := service.DB.Begin()
	helper.PanicIFError(err)
	defer helper.RollbackOrCommit(tx)
	HasPass, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	helper.PanicIFError(err)

	user := domain.User{
		Email:        request.Email,
		FirsName:     request.FirsName,
		LastName:     request.LastName,
		Password:     string(HasPass),
		ProfileImage: "https://asset.kompas.com/crops/BD5magRmejVwX0vqq28Ud2LxPXM=/0x1536:1920x2816/1200x800/data/photo/2023/04/28/644ba26ad81cb.jpg",
		Saldo:        0,
	}
	service.UserRepository.Save(ctx, tx, user)
	return nil
}

func (service *UserServiceImpl) Login(ctx context.Context, request web.UserLoginRequst) (web.UserResponse, error) {
	err := service.Validate.Struct(request)
	helper.PanicIFError(err)
	tx, err := service.DB.Begin()
	helper.PanicIFError(err)
	defer helper.RollbackOrCommit(tx)

	user, err := service.UserRepository.FindByEmail(ctx, tx, request.Email)
	if err != nil {
		return web.UserResponse{}, errors.New("invalid email")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return web.UserResponse{}, errors.New("password is not valid")
	}

	token, err := helper.GenerateJWT(user.Id, user.Email)
	return web.UserResponse{Token: "Bearer " + token}, err
}
