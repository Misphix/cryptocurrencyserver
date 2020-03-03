package apiprovider

// APIProvider is a interface define API provider should provide
// what we need
type APIProvider interface {
	GetLatestPrice(currency Currency) (float64, error)
}

// Currency it a enum define query currency
type Currency string

const (
	// Usd means USD
	Usd Currency = "usd"
	// Twd means TWD
	Twd = "twd"
)
