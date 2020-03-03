package querier

import (
	"errors"

	"github.com/misphix/cryptocurrencyserver/flowcontrol"

	"github.com/misphix/cryptocurrencyserver/apiprovider"
	"github.com/misphix/cryptocurrencyserver/configreader"
)

type providerResource struct {
	api           apiprovider.APIProvider
	lastTimePrice float64
	flowControl   *flowcontrol.FlowController
}

var config = configreader.ReadConfig()

var providerResources = map[string]providerResource{
	"CoinMarketCap": providerResource{
		api:           &apiprovider.CoinMarketCap{URL: apiprovider.CoinMarketCapURL, APIKey: config.CoinMarketCapKey},
		lastTimePrice: 0,
		flowControl:   flowcontrol.New(config.MaxSizeOfBucket, config.SecondPerToken),
	},
	"CryptoCompare": providerResource{
		api:           &apiprovider.CryptoComapre{URL: apiprovider.CryptoComapreURL, APIKey: config.CryptoCompareKey},
		lastTimePrice: 0,
		flowControl:   flowcontrol.New(config.MaxSizeOfBucket, config.SecondPerToken),
	},
	"CoinGecko": providerResource{
		api:           &apiprovider.CoinGecko{URL: apiprovider.CoinGeckoURL},
		lastTimePrice: 0,
		flowControl:   flowcontrol.New(config.MaxSizeOfBucket, config.SecondPerToken),
	},
}

// GetLatestPrice will get the latest price of specific provider.
// It will get last time's result if can't get the latest price
func GetLatestPrice(p string, currency apiprovider.Currency) (float64, error) {
	if p == "" {
		p = "CoinGecko"
	}

	provider, ok := providerResources[p]
	if ok {
		if provider.flowControl.AcquirePermission() {
			price, err := provider.api.GetLatestPrice(currency)

			if err == nil {
				provider.lastTimePrice = price
			}
		}

		return provider.lastTimePrice, nil
	}
	return 0, errors.New("Wrong provider parameter")
}

// AddTestProvider is a test only function.
// Do not use it in normal condition.
func AddTestProvider(name string, provider apiprovider.APIProvider) {
	providerResources[name] = providerResource{
		api:           provider,
		lastTimePrice: 0,
		flowControl:   flowcontrol.New(0, 0),
	}
}
