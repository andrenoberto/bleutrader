package services

type Response struct {
	Success bool   `json:"success,string"`
	Message string `json:"message"`
}
