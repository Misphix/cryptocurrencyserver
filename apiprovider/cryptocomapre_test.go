package apiprovider_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/misphix/cryptocurrencyserver/apiprovider"
)

type inputData struct {
	aPIKey   string
	currency apiprovider.Currency
}

func TestCryptoCompareGetLatestPrice(t *testing.T) {
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
			query := r.URL.Query()
			if len(query) != 2 {
				t.Errorf("Parameters number expected 2, actual %d", len(query))
			}

			if len(query["fsym"]) != 1 {
				t.Errorf("Parameter fsym number expected 1, actual %d", len(query["fsym"]))
			} else if query["fsym"][0] != "BTC" {
				t.Errorf("Parameter fsym value expected BTC, actual %q", query["fsym"][0])
			}

			if len(query["tsyms"]) != 1 {
				t.Errorf("Parameter tsyms number expected 1, actual %d", len(query["tsyms"]))
			} else if query["tsyms"][0] != string(c.in.currency) {
				t.Errorf("Parameter tsyms value expected %q, actual %q", string(c.in.currency), query["tsyms"][0])
			}

			result := map[string]float64{
				strings.ToUpper(string(c.in.currency)): c.want,
			}
			w.WriteHeader(200)
			bytes, _ := json.Marshal(result)
			w.Write(bytes)
		}))
		defer ts.Close()

		cryptoCompare := apiprovider.CryptoComapre{URL: ts.URL, APIKey: c.in.aPIKey}
		value, err := cryptoCompare.GetLatestPrice(c.in.currency)
		if err != nil {
			t.Errorf("GetLastestPrice error shold be empty, actual %q", err)
		}

		if value != c.want {
			t.Errorf("Expected %f, actual %f", c.want, value)
		}
	}
}

func TestCryptoCompareGetLatestPriceError(t *testing.T) {
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

		cryptoComare := apiprovider.CryptoComapre{URL: ts.URL}
		_, err := cryptoComare.GetLatestPrice(apiprovider.Usd)
		if err == nil {
			t.Errorf("It should be error %d", code)
		}
	}
}
