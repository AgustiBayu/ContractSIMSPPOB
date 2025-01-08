package web

type UserProfileUpdateImageRequest struct {
	Id           int    `json:"id" validate:"required"`
	ProfileImage string `json:"profile_image" validate:"required"`
}
