package main

import (
	"github.com/gin-gonic/gin"
	"github.com/misphix/cryptocurrencyserver/apiprovider"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1/cryptocurrency/")
	{
		v1.GET("/", queryPrice)
	}
	router.Run()
}

func queryPrice(context *gin.Context) {

	var currency apiprovider.Currency
	switch c := context.Query("currency"); c {
	case "twd":
		currency = apiprovider.Twd
	case "usd":
		fallthrough
	default:
		currency = apiprovider.Usd
	}

	var provider apiprovider.APIProvider
	switch p := context.Query("provider"); p {
	case "CoinMarketCap":
		token := "1cc823b9-41de-49ec-9f93-33d16ebf1860"
		provider = apiprovider.CoinMarketCap{APIKey: token}

	case "CoinGecko":
		fallthrough
	default:
		provider = apiprovider.CoinGecko{}
	}

	price := provider.GetLatestPrice(currency)
	context.JSON(200, gin.H{
		"BTC": price,
	})
}
