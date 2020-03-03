package main

import (
	"github.com/gin-gonic/gin"
	"github.com/misphix/cryptocurrencyserver/apiprovider"
	"github.com/misphix/cryptocurrencyserver/configreader"
	"github.com/misphix/cryptocurrencyserver/querier"
	"github.com/misphix/cryptocurrencyserver/usercontroller"
)

var currencies = map[string]apiprovider.Currency{
	"twd": apiprovider.Twd,
	"usd": apiprovider.Usd,
	"":    apiprovider.Usd,
}

func main() {
	config := configreader.ReadConfig()
	usercontroller.MaxTime = config.UserMaxQueryPerDay
	router := gin.Default()
	v1 := router.Group("/api/v1/cryptocurrency/")
	{
		v1.GET("/", queryPrice)
	}
	router.Run()
}

func queryPrice(context *gin.Context) {

	if !usercontroller.QuerryAcquire(context.ClientIP()) {
		context.JSON(200, gin.H{
			"error": "Exceed 24 hours limit",
		})
		return
	}

	currency, ok := currencies[context.Query("currency")]
	if !ok {
		context.JSON(200, gin.H{
			"error": "Wrong currency parameter",
		})
		return
	}

	price, err := querier.GetLatestPrice(context.Query("provider"), currency)

	if err != nil {
		context.JSON(200, gin.H{
			"error": "Wrong provider parameter",
		})
		return
	}

	context.JSON(200, gin.H{
		"BTC": price,
	})
}
