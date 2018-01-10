package entity

type OrderBook struct {
	Buy  []Order `json:"buy"`
	Sell []Order `json:"sell"`
}
