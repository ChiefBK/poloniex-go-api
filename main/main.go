package main

import (
	"poloniex-go-api"
	"fmt"
)

func main(){
	poloniex := poloniex_go_api.New("TIDAAD1L-YHIL9FNO-VNLI7T9W-WYYLZWRI", "6443cebc72db9d1deb6fa4d5913db8f65dffc448a4700610889e67a2ca0bf24f74aef3294bc25fc122dfc1df40921ad7474df51b47a8c806d35eb33d17a6c593")

	fmt.Println("returning open loan offers")
	resp := poloniex.ReturnOpenLoanOffers()
	poloniex_go_api.PrintResponse(resp)

	fmt.Println("returning ticker")
	resp2 := poloniex.ReturnTicker()
	poloniex_go_api.PrintResponse(resp2)
}
