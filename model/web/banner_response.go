package web

type BannerResponse struct {
	Id          int    `json:"id"`
	BannerName  string `json:"banner_name"`
	BannerImage string `json:"banner_image"`
	Description string `json:"description"`
}
