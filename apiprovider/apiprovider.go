package apiprovider

// APIProvider is a interface define API provider should provide
// what we need
type APIProvider interface {
	GetLatestPrice(currency Currency) float32
}

type Currency string

const (
	Usd Currency = "usd"
	Twd          = "twd"
)
