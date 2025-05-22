package dto

type WebResponse[T any] struct {
	Message string         `json:"message,omitempty"`
	Data    T              `json:"data,omitempty"`
	Errors  *ErrorResponse `json:"errors,omitempty"`
}

type ErrorResponse struct {
	Message string `json:"message,omitempty"`
}

type WebResponseSwagger struct {
	Message string          `json:"message"`
	Data    ShortenResponse `json:"data"`
}

type ErrorResponseSwagger struct {
    Message string `json:"errors"`
}


