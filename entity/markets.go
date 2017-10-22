package entity

type Markets struct {
	MarketCurrency     string  `json:"MarketCurrency"`
	BaseCurrency       string  `json:"BaseCurrency"`
	MarketCurrencyLong string  `json:"MarketCurrencyLong"`
	BaseCurrencyLong   string  `json:"BaseCurrencyLong"`
	MinTradeSize       float32 `json:"MinTradeSize"`
	MarketName         string  `json:"MarketName"`
	IsActive           bool    `json:"IsActive"`
	Created            string  `json:"Created"`
	Notice             string  `json:"Notice"`
	IsSponsored        bool    `json:"IsSponsored"`
	LogoUrl            string  `json:"LogoUrl"`
}
