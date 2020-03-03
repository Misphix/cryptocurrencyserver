package querier_test

import (
	"errors"
	"testing"

	"github.com/misphix/cryptocurrencyserver/querier"

	"github.com/misphix/cryptocurrencyserver/apiprovider"
)

const (
	ExpectedValue = 963.258
)

var counter = 0

type TestProvider struct{}

func (t TestProvider) GetLatestPrice(currency apiprovider.Currency) (float64, error) {
	if counter == 0 {
		return ExpectedValue, nil
	}

	return 0, errors.New("")
}

func TestQuerierGetLastTimePrice(t *testing.T) {
	querier.AddTestProvider("9bf29fb9-b45d-4c6a-8b04-c222bf53cd66", TestProvider{})

	price, _ := querier.GetLatestPrice("9bf29fb9-b45d-4c6a-8b04-c222bf53cd66", apiprovider.Usd)
	if price != ExpectedValue {
		t.Errorf("First query expected %g, actual %g", ExpectedValue, price)
	}

	price, _ = querier.GetLatestPrice("9bf29fb9-b45d-4c6a-8b04-c222bf53cd66", apiprovider.Usd)
	if price != ExpectedValue {
		t.Errorf("Second query expected %g, actual %g", ExpectedValue, price)
	}
}
