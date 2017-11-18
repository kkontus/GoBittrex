package entity

type OpenOrdersResponse struct {
	ResponseGeneric
	Result []OpenOrders `json:"result"`
}
