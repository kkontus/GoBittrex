package entity

type TicksResponse struct {
	ResponseGeneric
	Result []Historical `json:"result"`
}
