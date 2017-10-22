package entity

type CurrenciesResponse struct {
	ResponseGeneric
	Result []Currencies `json:"result"`
}
