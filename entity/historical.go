package entity

type Historical struct {
	Open       float32 `json:"O"`
	High       float32 `json:"H"`
	Low        float32 `json:"L"`
	Close      float32 `json:"C"`
	Volume     float32 `json:"V"`
	Timestamp  string  `json:"T"`
	BaseVolume float32 `json:"BV"`
}

type TicksResponse struct {
	ResponseGeneric
	Result []Historical `json:"result"`
}
