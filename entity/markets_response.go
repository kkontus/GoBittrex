package entity

type MarketsResponse struct {
	ResponseGeneric
	Result []Markets `json:"result"`
}
