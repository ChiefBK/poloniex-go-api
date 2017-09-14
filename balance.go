package poloniex_go_api

type Balance struct {
	Available string `json:"available"`
	OnOrders  string `json:"onOrders"`
	BtcValue  string `json:"btcValue"`
}
