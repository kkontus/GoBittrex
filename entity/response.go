package entity

type ResponseGeneric struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
