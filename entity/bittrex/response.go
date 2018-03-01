package bittrex

type ResponseGeneric struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
