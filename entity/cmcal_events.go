package entity

type CmcalEvents struct {
	Id                int        `json:"id"`
	Title             string     `json:"title"`
	Coins             []CoinMeta `json:"coins"`
	DateEvent         string     `json:"date_event"`
	CreatedDate       string     `json:"created_date"`
	Description       string     `json:"description"`
	Proof             string     `json:"proof"`
	Source            string     `json:"source"`
	IsHot             bool       `json:"is_hot"`
	VoteCount         int        `json:"vote_count"`
	PositiveVoteCount int        `json:"positive_vote_count"`
	Percentage        int        `json:"percentage"`
	Categories        []string   `json:"categories"`
	TipSymbol         string     `json:"tip_symbol"`
	TipAddress        string     `json:"tip_adress"`
	TwitterAccount    string     `json:"twitter_account"`
	CanOccurBefore    bool       `json:"can_occur_before"`
}

type CoinMeta struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}
