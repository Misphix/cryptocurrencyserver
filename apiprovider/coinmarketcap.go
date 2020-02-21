package apiprovider

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

// CoinMarketCap will set the basic data of CoinMarketCap needs
type CoinMarketCap struct {
	APIKey string
}

// Quote represent cryptocurreny's value of a specific currency
type Quote struct {
	Price     float32
	Volume24h float32 `json:"volume_24h"`
}

// Data represent a cryptocurrency status
type Data struct {
	ID    int
	Name  string
	Quote map[string]Quote
}

// Status represent the status of response
type Status struct {
	Timestamp   time.Time
	Elapsed     int
	CreditCount int       `json:"credit_count"`
	LastUpdated time.Time `json:"last_updated"`
}

// Response is CoinMarketCap's quotes reponse
type Response struct {
	Data   map[string]Data
	Status Status
}

// GetLatestPrice will get latest BTC price with USD
func (c *CoinMarketCap) GetLatestPrice(currency Currency) float32 {
	currencyID := map[Currency]int{
		Usd: 2781,
		Twd: 2811,
	}[currency]

	client := &http.Client{}
	request, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := url.Values{}
	q.Add("id", "1")
	q.Add("convert_id", strconv.Itoa(currencyID))

	request.Header.Set("Accepts", "application/json")
	request.Header.Add("X-CMC_PRO_API_KEY", c.APIKey)
	request.URL.RawQuery = q.Encode()

	r, err := client.Do(request)
	if err != nil {
		log.Print(err)
		// TODO error handling
	}

	var response Response
	err = json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		log.Print(err)
		// TODO error handling
	}

	return response.Data["1"].Quote[strconv.Itoa(currencyID)].Price
}
