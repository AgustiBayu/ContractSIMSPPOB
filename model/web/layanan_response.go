package web

type LayananResponse struct {
	Id           int    `json:"id"`
	ServiceCode  string `json:"service_code"`
	ServiceName  string `json:"service_name"`
	ServiceIcon  string `json:"service_icon"`
	ServiceTarif int    `json:"service_tarif"`
}
