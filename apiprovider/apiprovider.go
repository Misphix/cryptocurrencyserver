package apiprovider

// APIProvider is a interface define API provider should provide
// what we need
type APIProvider interface {
	GetLatestPrice(currency Currency) (float64, error)
}

type Currency string

const (
	Usd Currency = "usd"
	Twd          = "twd"
)
