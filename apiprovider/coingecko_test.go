package apiprovider_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/misphix/cryptocurrencyserver/apiprovider"
)

func TestCoinGeckoGetLatestPrice(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := r.URL
		if len(url.Query()) != 2 {
			t.Errorf("Parameters number expected 2, actual %d", len(url.Query()))
		}
	}))
	defer ts.Close()

	coinGecko := apiprovider.CoinGecko{URL: ts.URL}
	coinGecko.GetLatestPrice(apiprovider.Usd)
}
