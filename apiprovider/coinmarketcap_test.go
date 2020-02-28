package apiprovider_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/misphix/cryptocurrencyserver/apiprovider"
)

func TestCoinMarketCapLatestPriceUsd(t *testing.T) {
	currencyIDMap := map[apiprovider.Currency]int{
		apiprovider.Usd: 2781,
		apiprovider.Twd: 2811,
	}

	type inputData struct {
		aPIKey   string
		currency apiprovider.Currency
	}

	cases := []struct {
		in   inputData
		want float64
	}{
		{inputData{"c00a16de-55a7-421c-98be-45ba96bc94c5", apiprovider.Usd}, 3.1415926},
		{inputData{"6024f3d6-0969-4940-acdf-b6f29c7a2043", apiprovider.Usd}, 2.7182818},
		{inputData{"9cbbea16-3502-437c-b7ca-59311a679c11", apiprovider.Twd}, 81000},
		{inputData{"0b90b3d6-2c2f-4ee2-bd73-f83bb396da8c", apiprovider.Twd}, 9487},
	}

	for _, c := range cases {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("X-CMC_PRO_API_KEY") != c.in.aPIKey {
				t.Errorf("API key not fit. Expected: %q, acutal: %q", c.in.aPIKey, r.Header.Get("X-CMC_PRO_API_KEY"))
			}

			query := r.URL.Query()
			if len(query) != 2 {
				t.Errorf("Parameters number expected 2, actual %d", len(query))
			}

			if len(query["id"]) != 1 {
				t.Errorf("Parameter ids number expected 1, actual %d", len(query["ids"]))
			} else if query["id"][0] != "1" {
				t.Errorf("Parameter ids value expected 1, actual %q", query["ids"][0])
			}

			if len(query["convert_id"]) != 1 {
				t.Errorf("Parameter vs_currencies number expected 1, actual %d", len(query["ids"]))
			} else if query["convert_id"][0] != strconv.Itoa(currencyIDMap[c.in.currency]) {
				t.Errorf("Parameter ids value expected %d, actual %q", currencyIDMap[c.in.currency], query["vs_currencies"][0])
			}

			result := `{
"data": {
"1": {
"id": 1,
"name": "Bitcoin",
"quote": {
"%d": {
"price": %g,
"volume_24h": 4314444687.5194,
"last_updated": "2018-08-09T21:56:28.000Z"
}
}
}
},
"status": {
"timestamp": "2020-02-28T06:48:22.528Z",
"elapsed": 10,
"credit_count": 1
}
}`
			result = fmt.Sprintf(result, currencyIDMap[c.in.currency], c.want)
			w.WriteHeader(200)
			w.Write([]byte(result))
		}))
		defer ts.Close()

		coinMarketCap := apiprovider.CoinMarketCap{APIKey: c.in.aPIKey, URL: ts.URL}
		value, err := coinMarketCap.GetLatestPrice(c.in.currency)
		if err != nil {
			t.Errorf("GetLastestPrice error shold be empty, actual %q", err)
		}

		if value != c.want {
			t.Errorf("Expected %f, actual %f", c.want, value)
		}
	}
}

func TestCoinMarketCapLatestPriceUsdError(t *testing.T) {
	errorCodes := []int{
		400,
		401,
		403,
		404,
		405,
		500,
		501,
	}

	for _, code := range errorCodes {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(code)
		}))
		defer ts.Close()

		coinMarketCap := apiprovider.CoinMarketCap{APIKey: "2073799d-5561-48e7-abef-c9f48f88695c", URL: ts.URL}
		_, err := coinMarketCap.GetLatestPrice(apiprovider.Usd)
		if err == nil {
			t.Errorf("It should be error %d", code)
		}
	}
}
