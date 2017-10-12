package main

import (
	"log"
	"os"
	"poloniex-go-api"
)

func main() {
	apiKey := os.Getenv("POLONIEX_API_KEY")
	apiSecret := os.Getenv("POLONIEX_API_SECRET")

	command := "returnCompleteBalances"
	method := "POST"
	data := make(map[string]string)
	//data["account"] = "all"

	poloniex := poloniex_go_api.New(apiKey, apiSecret)

	response, err := poloniex.Client.Do(method, command, data)

	if err != nil {
		log.Println("there was an error making the request")
		return
	}

	log.Printf("RESPONSE: %s", response)
}
