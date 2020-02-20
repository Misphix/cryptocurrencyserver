package apiprovider

// APIProvider is a interface define API provider should provide
// what we need
type APIProvider interface {
	GetLatestPrice() float32
}
