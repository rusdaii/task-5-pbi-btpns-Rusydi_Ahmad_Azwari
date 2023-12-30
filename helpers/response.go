package helpers

type SuccesResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type LoginResponse struct {
	Success     bool        `json:"success"`
	Message     string      `json:"message"`
	AccessToken string      `json:"accessToken"`
	Data        interface{} `json:"data"`
}
