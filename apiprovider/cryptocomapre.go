package apiprovider

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

// CryptoComapreURL is url of CoinMarketCap
const CryptoComapreURL = "https://min-api.cryptocompare.com/data/price"

// CryptoComapre will set the basic data of CoinMarketCap needs
type CryptoComapre struct {
	URL    string
	APIKey string
}

// GetLatestPrice will get latest BTC price with USD
func (c *CryptoComapre) GetLatestPrice(currency Currency) (float64, error) {
	request, err := http.NewRequest("GET", c.URL, nil)
	if err != nil {
		return 0, err
	}

	q := url.Values{}
	q.Add("fsym", "BTC")
	q.Add("tsyms", string(currency))

	request.Header.Set("Accepts", "application/json")
	request.Header.Set("Apikey", c.APIKey)
	request.URL.RawQuery = q.Encode()

	client := &http.Client{}
	r, err := client.Do(request)
	if err != nil {
		return 0, err
	}

	var response map[string]float64
	err = json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		return 0, err
	}

	return response[strings.ToUpper(string(currency))], nil
}
