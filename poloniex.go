package poloniex_go_api

import (
	"encoding/json"
	"fmt"
	"time"
	"log"
)

func New(apiKey, apiSecret string) *Poloniex {
	client := NewClient(apiKey, apiSecret)
	return &Poloniex{client}
}

type Poloniex struct {
	Client *Client
}

func (p *Poloniex) ReturnTicker() (resp *ReturnTickerResponse) {
	clientRes, _ := p.Client.do("GET", "returnTicker", nil)

	resp = new(ReturnTickerResponse)
	err := json.Unmarshal(clientRes, &resp.Response)
	//err := getJson(res, tickerJsonRes)

	if err == nil {
		resp.Err = nil
		return
	} else {
		fmt.Println("error on unmarshal")
		resp.Err = err
		return
	}
}

func (p *Poloniex) ReturnOpenLoanOffers() (resp *ReturnOpenLoanOffersResponse) {
	clientRes, clientResError := p.Client.do("POST", "returnOpenLoanOffers", nil)

	if clientResError != nil {
		fmt.Println("error making request (do)")
	}

	resp = new(ReturnOpenLoanOffersResponse)
	err := json.Unmarshal(clientRes, &resp.Response)

	if err == nil {
		resp.Err = nil
		return
	} else {
		resp.Err = err
		return
	}
}

func (p *Poloniex) ReturnLoanOrders(respCh chan *ReturnLoanOrdersResponse) {
	defer close(respCh)

	data := make(map[string]string)
	data["currency"] = "BTC"

	res, err := p.Client.do("GET", "returnLoanOrders", data)

	returnLoanOrders := new(ReturnLoanOrdersResponse)

	if err != nil {
		returnLoanOrders.Response = nil
		returnLoanOrders.Err = err
		respCh <- returnLoanOrders
		return
	}

	err = json.Unmarshal(res, &returnLoanOrders.Response)

	if err != nil {
		returnLoanOrders.Err = err
		returnLoanOrders.Response = nil
		respCh <- returnLoanOrders
		return
	}

	returnLoanOrders.Err = nil
	respCh <- returnLoanOrders
}

func (p *Poloniex) ReturnLoanOffers(offersCh chan []*Order) {
	defer close(offersCh)

	returnLoanOrdersCh := make(chan *ReturnLoanOrdersResponse)
	go p.ReturnLoanOrders(returnLoanOrdersCh)
	loanOrders := <-returnLoanOrdersCh

	if loanOrders.Err != nil {
		offersCh <- nil
	}

	offersCh <- loanOrders.Response["offers"]
}

func (p *Poloniex) ReturnBalances(respCh chan *ReturnBalancesResponse) {
	defer close(respCh)

	res, err := p.Client.do("POST", "returnBalances", nil)

	balancesResp := new(ReturnBalancesResponse)

	if err != nil {
		fmt.Println("error making request (do)")
		balancesResp.Err = err
		balancesResp.Response = nil
		respCh <- balancesResp
		return
	}

	err = json.Unmarshal(res, &balancesResp.Response)

	if err != nil {
		fmt.Println("error unmarshalling")
		balancesResp.Err = err
		balancesResp.Response = nil
		respCh <- balancesResp
		return
	}

	balancesResp.Err = nil
	respCh <- balancesResp
}

func (p *Poloniex) ReturnCompleteBalancesBtc(balanceCh chan *Balance) {
	defer close(balanceCh)

	completeBalancesCh := make(chan *ReturnCompleteBalancesResponse)

	go p.ReturnCompleteBalances(completeBalancesCh)

	completeBalances := <-completeBalancesCh

	balance := new(Balance)
	balance.Available = completeBalances.Response["BTC"].Available
	balance.BtcValue = completeBalances.Response["BTC"].BtcValue
	balance.OnOrders = completeBalances.Response["BTC"].OnOrders
	balanceCh <- balance
}

func (p *Poloniex) ReturnCompleteBalances(respCh chan *ReturnCompleteBalancesResponse) {
	data := make(map[string]string)
	data["account"] = "all"
	clientRes, clientResError := p.Client.do("POST", "returnCompleteBalances", data)

	if clientResError != nil {
		fmt.Println("error making request (do)")
	}

	resp := new(ReturnCompleteBalancesResponse)
	err := json.Unmarshal(clientRes, &resp.Response)

	if err == nil {
		resp.Err = nil
		respCh <- resp
	} else {
		resp.Err = err
		respCh <- resp
	}
}

func (p *Poloniex) ReturnLendingHistory(start, end time.Time, respCh chan *ReturnLendingHistoryResponse) {
	defer close(respCh)

	data := make(map[string]string)
	data["start"] = fmt.Sprintf("%d", start.Unix())
	data["end"] = fmt.Sprintf("%d", end.Unix())
	res, err := p.Client.do("POST", "returnLendingHistory", data)

	lendingHistoryResponse := new(ReturnLendingHistoryResponse)

	if err != nil {
		lendingHistoryResponse.Response = nil
		lendingHistoryResponse.Err = err
		respCh <- lendingHistoryResponse
		return
	}

	json.Unmarshal(res, &lendingHistoryResponse.Response)
	lendingHistoryResponse.Err = nil
	respCh <- lendingHistoryResponse
}

func (p *Poloniex) ReturnActiveLoans(respCh chan *ReturnActiveLoansResponse) {
	defer close(respCh)

	res, err := p.Client.do("POST", "returnActiveLoans", nil)

	activeLoansResponse := new(ReturnActiveLoansResponse)

	if err != nil {
		activeLoansResponse.Response = nil
		activeLoansResponse.Err = err
		respCh <- activeLoansResponse
		return
	}

	json.Unmarshal(res, &activeLoansResponse.Response)
	activeLoansResponse.Err = nil
	respCh <- activeLoansResponse
}

func (p *Poloniex) ReturnChartData(currencyPair string, period int, start, end time.Time) (chartDataResponse *ReturnChartDataResponse) {
	log.Println("Returning Chart Data")

	data := make(map[string]string)
	data["currencyPair"] = currencyPair
	data["period"] = fmt.Sprintf("%d", period)
	data["start"] = fmt.Sprintf("%d", start.Unix())
	data["end"] = fmt.Sprintf("%d", end.Unix())

	res, err := p.Client.do("GET", "returnChartData", data)

	chartDataResponse = new(ReturnChartDataResponse)

	if err != nil {
		chartDataResponse.Response = nil
		chartDataResponse.Err = err
		return
	}

	json.Unmarshal(res, &chartDataResponse.Response)
	chartDataResponse.Err = nil
	return
}

func PrintResponse(resp interface{}) {
	j, _ := json.Marshal(resp)
	fmt.Println(string(j))
}
