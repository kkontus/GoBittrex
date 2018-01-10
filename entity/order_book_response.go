package entity

type OrderBookResponse struct {
	ResponseGeneric
	Result OrderBook `json:"result"`
}
