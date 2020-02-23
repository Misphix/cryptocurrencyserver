package apiprovider

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

// CoinGecko will set the basic data of CoinGecko needs
type CoinGecko struct {
}

// GetLatestPrice will get latest BTC price with USD
func (c CoinGecko) GetLatestPrice(currency Currency) float32 {
	client := &http.Client{}
	request, err := http.NewRequest("GET", "https://api.coingecko.com/api/v3/simple/price", nil)
	if err != nil {
		log.Print(err)
		// TODO error handling
	}

	q := url.Values{}
	q.Add("ids", "bitcoin")
	q.Add("vs_currencies", string(currency))

	request.Header.Set("Accepts", "application/json")
	request.URL.RawQuery = q.Encode()

	r, err := client.Do(request)
	if err != nil {
		log.Print(err)
		// TODO error handling
	}

	var response map[string]map[string]float32
	json.NewDecoder(r.Body).Decode(&response)

	return response["bitcoin"][string(currency)]
}
