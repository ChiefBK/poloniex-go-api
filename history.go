package poloniex_go_api

type History struct {
	ID       int    `json:"id"`
	Currency string `json:"currency"`
	Rate     string `json:"rate"`
	Amount   string `json:"amount"`
	Duration string `json:"duration"`
	Interest string `json:"interest"`
	Fee      string `json:"fee"`
	Earned   string `json:"earned"`
	Open     string `json:"open"`
	Close    string `json:"close"`
}
