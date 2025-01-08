package web

type UserProfileUpdateRequest struct {
	Id       int    `json:"id" validate:"required"`
	FirsName string `json:"firs_name" validate:"required"`
	LastName string `json:"last_name" validate:"required"`
}
