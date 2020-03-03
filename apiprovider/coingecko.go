package apiprovider

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

// CoinGeckoURL is url of CoinGecko
const CoinGeckoURL = "https://api.coingecko.com/api/v3/simple/price"

// CoinGecko will set the basic data of CoinGecko needs
type CoinGecko struct {
	URL string
}

// GetLatestPrice will get latest BTC price with USD
func (c *CoinGecko) GetLatestPrice(currency Currency) (float64, error) {
	request, err := http.NewRequest("GET", c.URL, nil)
	if err != nil {
		return 0, err
	}

	q := url.Values{}
	q.Add("ids", "bitcoin")
	q.Add("vs_currencies", string(currency))

	request.Header.Set("Accepts", "application/json")
	request.URL.RawQuery = q.Encode()

	client := &http.Client{}
	r, err := client.Do(request)
	if err != nil {
		return 0, err
	}

	if r.StatusCode != 200 {
		return 0, errors.New(r.Status)
	}

	var response map[string]map[string]float64
	json.NewDecoder(r.Body).Decode(&response)

	return response["bitcoin"][string(currency)], nil
}
