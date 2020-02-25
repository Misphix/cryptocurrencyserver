package main

import (
	"github.com/gin-gonic/gin"
	"github.com/misphix/cryptocurrencyserver/apiprovider"
	"github.com/misphix/cryptocurrencyserver/configreader"
)

var providers = make(map[string]apiprovider.APIProvider)
var currencies = make(map[string]apiprovider.Currency)

func main() {
	config := configreader.ReadConfig()
	initializeParametersMap(config)

	router := gin.Default()
	v1 := router.Group("/api/v1/cryptocurrency/")
	{
		v1.GET("/", queryPrice)
	}
	router.Run()
}

func initializeParametersMap(config configreader.Config) {
	// Initialize provider map
	providers["CoinMarketCap"] = apiprovider.CoinMarketCap{APIKey: config.CoinMarketCapKey}
	coinGecko := apiprovider.CoinGecko{}
	providers["CoinGecko"] = coinGecko
	providers[""] = coinGecko

	// Initialize currencies map
	currencies["twd"] = apiprovider.Twd
	currencies["usd"] = apiprovider.Usd
	currencies[""] = apiprovider.Usd
}

func queryPrice(context *gin.Context) {
	currency, ok := currencies[context.Query("currency")]
	if !ok {
		context.JSON(200, gin.H{
			"error": "Wrong currency parameter",
		})
		return
	}

	provider, ok := providers[context.Query("provider")]
	if ok {
		price := provider.GetLatestPrice(currency)
		context.JSON(200, gin.H{
			"BTC": price,
		})
	} else {
		context.JSON(200, gin.H{
			"error": "Wrong provider parameter",
		})
		return
	}
}
