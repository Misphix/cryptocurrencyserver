package apiprovider_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/misphix/cryptocurrencyserver/apiprovider"
)

func TestCoinGeckoGetLatestPriceUsd(t *testing.T) {
	cases := []struct {
		in   apiprovider.Currency
		want float64
	}{
		{apiprovider.Usd, 3.1415926},
		{apiprovider.Usd, 2.7182818},
		{apiprovider.Twd, 81000},
		{apiprovider.Usd, 9487},
	}

	for _, c := range cases {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			query := r.URL.Query()
			if len(query) != 2 {
				t.Errorf("Parameters number expected 2, actual %d", len(query))
			}

			if len(query["ids"]) != 1 {
				t.Errorf("Parameter ids number expected 1, actual %d", len(query["ids"]))
			} else if query["ids"][0] != "bitcoin" {
				t.Errorf("Parameter ids value expected bitcoin, actual %q", query["ids"][0])
			}

			if len(query["vs_currencies"]) != 1 {
				t.Errorf("Parameter vs_currencies number expected 1, actual %d", len(query["vs_currencies"]))
			} else if query["vs_currencies"][0] != string(c.in) {
				t.Errorf("Parameter vs_currencies value expected %q, actual %q", string(c.in), query["vs_currencies"][0])
			}

			result := map[string]map[string]float64{
				"bitcoin": map[string]float64{
					string(c.in): c.want,
				},
			}
			w.WriteHeader(200)
			bytes, _ := json.Marshal(result)
			w.Write(bytes)
		}))
		defer ts.Close()

		coinGecko := apiprovider.CoinGecko{URL: ts.URL}
		value, err := coinGecko.GetLatestPrice(c.in)
		if err != nil {
			t.Errorf("GetLastestPrice error shold be empty, actual %q", err)
		}

		if value != c.want {
			t.Errorf("Expected %f, actual %f", c.want, value)
		}
	}
}

func TestCoinGeckoGetLatestPriceUsdError(t *testing.T) {
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

		coinGecko := apiprovider.CoinGecko{URL: ts.URL}
		_, err := coinGecko.GetLatestPrice(apiprovider.Usd)
		if err == nil {
			t.Errorf("It should be error %d", code)
		}
	}
}
