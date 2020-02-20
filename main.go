package main

import (
	"fmt"

	"github.com/misphix/cryptocurrencyserver/apiprovider"
)

func main() {
	token := "1cc823b9-41de-49ec-9f93-33d16ebf1860"
	coinMarketCap := apiprovider.CoinMarketCap{APIKey: token}
	fmt.Println(coinMarketCap.GetLatestPrice())

	coinGecko := apiprovider.CoinGecko{}
	fmt.Println(coinGecko.GetLatestPrice())
}
