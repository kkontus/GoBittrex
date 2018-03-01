package bittrex

type OrderBook struct {
	Buy  []Order `json:"buy"`
	Sell []Order `json:"sell"`
}

type Order struct {
	Quantity float32 `json:"Quantity"`
	Rate     float32 `json:"Rate"`
}

type OrderBookResponse struct {
	ResponseGeneric
	Result OrderBook `json:"result"`
}
