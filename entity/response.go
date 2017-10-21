package entity

type Response struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Result  []BittrexResponse `json:"result"`
}
