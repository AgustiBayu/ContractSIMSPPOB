package web

type UserRegisterRequst struct {
	Email    string `json:"email" validate:"required"`
	FirsName string `json:"firs_name" validate:"required"`
	LastName string `json:"last_name" validate:"required"`
	Password string `json:"password" validate:"required"`
}
