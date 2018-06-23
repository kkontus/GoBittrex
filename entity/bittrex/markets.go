package bittrex

// Markets
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

type MarketsResponse struct {
	ResponseGeneric
	Result []Markets `json:"result"`
}

// Market Summaries
type MarketSummaries struct {
	MarketSummary
}

type MarketSummariesResponse struct {
	ResponseGeneric
	Result []MarketSummaries `json:"result"`
}

// Market Summary
type MarketSummary struct {
	MarketName        string  `json:"MarketName"`
	High              float32 `json:"High"`
	Low               float32 `json:"Low"`
	Volume            float32 `json:"Volume"`
	Last              float32 `json:"Last"`
	BaseVolume        float32 `json:"BaseVolume"`
	TimeStamp         string  `json:"TimeStamp"`
	Bid               float32 `json:"Bid"`
	Ask               float32 `json:"Ask"`
	OpenBuyOrders     int     `json:"OpenBuyOrders"`
	OpenSellOrders    int     `json:"OpenSellOrders"`
	PrevDay           float32 `json:"PrevDay"`
	Created           string  `json:"Created"`
	DisplayMarketName string  `json:"DisplayMarketName"`
}

type MarketSummaryResponse struct {
	ResponseGeneric
	Result []MarketSummary `json:"result"`
}
