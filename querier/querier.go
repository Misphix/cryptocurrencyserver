package querier

import (
	"errors"

	"github.com/misphix/cryptocurrencyserver/apiprovider"
	"github.com/misphix/cryptocurrencyserver/configreader"
)

var config = configreader.ReadConfig()
var providers = map[string]apiprovider.APIProvider{
	"CoinMarketCap": apiprovider.CoinMarketCap{URL: apiprovider.CoinMarketCapURL, APIKey: config.CoinMarketCapKey},
	"CryptoCompare": apiprovider.CryptoComapre{URL: apiprovider.CryptoComapreURL, APIKey: config.CryptoCompareKey},
	"CoinGecko":     apiprovider.CoinGecko{URL: apiprovider.CoinGeckoURL},
}

var lastTimePrice = map[string]float64{
	"CoinMarketCap": 0,
	"CryptoCompare": 0,
	"CoinGecko":     0,
}

// GetLatestPrice will get the latest price of specific provider.
// It will get last time's result if can't get the latest price
func GetLatestPrice(p string, currency apiprovider.Currency) (float64, error) {
	if p == "" {
		p = "CoinGecko"
	}

	provider, ok := providers[p]
	if ok {
		price, err := provider.GetLatestPrice(currency)

		if err != nil {
			return getLastTimePrice(p), nil
		}

		lastTimePrice[p] = price
		return price, nil
	}
	return 0, errors.New("Wrong provider parameter")
}

// AddTestProvider is a test only function.
// Do not use it in normal condition.
func AddTestProvider(name string, provider apiprovider.APIProvider) {
	providers[name] = provider
	lastTimePrice[name] = 0
}

func getLastTimePrice(p string) float64 {
	price, ok := lastTimePrice[p]
	if ok {
		return price
	}

	return 0
}
