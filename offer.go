package poloniex_go_api

type Offer struct {
	ID        int    `json:"id"`
	Rate      string `json:"rate"`
	Amount    string `json:"amount"`
	Duration  int    `json:"duration"`
	AutoRenew int    `json:"autoRenew"`
	Date      string `json:"date"`
}
