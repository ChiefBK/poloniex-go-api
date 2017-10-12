package poloniex_go_api

/*
	PUBLIC API RESPONSES
*/

type ReturnTickerResponse struct {
	Response map[string]Pair
	Err      error
}

type Return24VolumeResponse struct {
}

type ReturnOrderBookResponse struct {
}

type ReturnChartDataResponse struct {
	Response []*Candle
	Err      error
}

type ReturnCurrenciesResponse struct {
}

type ReturnLoanOrdersResponse struct {
	Response map[string][]*Order
	Err      error
}

/*
	TRADING API RESPONSES
*/

type ReturnBalancesResponse struct {
	Response map[string]string
	Err      error
}

type ReturnCompleteBalancesResponse struct {
	Data map[string]*Balance
	Err  error
}

type ReturnDepositAddressesResponse struct {
}

type GenerateNewAddressResponse struct {
}

type ReturnDepositsWithdrawalsResponse struct {
}

type ReturnOpenOrdersResponse struct {
}

type ReturnOrderTradesResponse struct {
}

/*
	........
*/

type ReturnAvailableAccountBalancesResponse struct {
	Response struct {
		exchange map[string]string
		margin   map[string]string
		lending  map[string]string
	}
	Err error
}

type TransferBalanceResponse struct {
	Response struct {
		success string
		message string
	}
	Err error
}

type CreateLoanOfferResponse struct {
	Response struct {
		success string
		message string
	}
	Err error
}

type CancelLoanOfferResponse struct {
	Response struct {
		success string
		message string
	}
	Err error
}

type ReturnOpenLoanOffersResponse struct {
	Response map[string][]*Offer
	Err      error
}

type ReturnActiveLoansResponse struct {
	Response map[string][]*Loan
	Err      error
}

type ReturnLendingHistoryResponse struct {
	Response []History
	Err      error
}

type ToggleAutoRenewResponse struct {
	Response struct {
		success int
		message int
	}
	Err error
}
