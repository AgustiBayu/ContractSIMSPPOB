package web

type UserProfileResponse struct {
	Id           int    `json:"id" validate:"required"`
	Email        string `json:"email" validate:"required"`
	FirsName     string `json:"firs_name" validate:"required"`
	LastName     string `json:"last_name" validate:"required"`
	ProfileImage string `json:"profile_image" validate:"required"`
}
