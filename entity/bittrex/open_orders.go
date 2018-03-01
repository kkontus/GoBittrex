package bittrex

type OpenOrders struct {
	OrderUuid         string  `json:"OrderUuid"`
	Exchange          string  `json:"Exchange"`
	OrderType         string  `json:"OrderType"`
	Quantity          float32 `json:"Quantity"`
	QuantityRemaining float32 `json:"QuantityRemaining"`
	Limit             float32 `json:"Limit"`
	CommissionPaid    float32 `json:"CommissionPaid"`
	Price             float32 `json:"Price"`
	PricePerUnit      float32 `json:"PricePerUnit"`
	Opened            string  `json:"Opened"`
	Closed            string  `json:"Closed"`
	CancelInitiated   bool    `json:"CancelInitiated"`
	ImmediateOrCancel bool    `json:"ImmediateOrCancel"`
	IsConditional     bool    `json:"IsConditional"`
	Condition         string  `json:"Condition"`
	ConditionTarget   string  `json:"ConditionTarget"`
}

type OpenOrdersResponse struct {
	ResponseGeneric
	Result []OpenOrders `json:"result"`
}
