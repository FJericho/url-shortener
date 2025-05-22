package dto

type ShortenRequest struct {
	Original  string `json:"original" validate:"required,url"`
}

type ShortenResponse struct {
	Original  string `json:"original" validate:"required,url"`
	ShortCode string `json:"short_code" validate:"omitempty,alphanum"`
}
