package poloniex_go_api

type Loan struct {
	ID        int    `json:"id"`
	Currency  string `json:"currency"`
	Rate      string `json:"rate"`
	Amount    string `json:"amount"`
	Range     int    `json:"range"`
	AutoRenew int    `json:"autoRenew"`
	Date      string `json:"date"`
	Fees      string `json:"fees"`
}
