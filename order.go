package poloniex_go_api

type Order struct {
	Rate     string `json:"rate"`
	Amount   string `json:"amount"`
	RangeMin int    `json:"rangeMin"`
	RangeMax int    `json:"rangeMax"`
}
