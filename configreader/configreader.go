package configreader

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config is a structure store config value
type Config struct {
	CoinMarketCapKey string
}

// ReadConfig will read config from current directory
func ReadConfig() Config {
	viper.SetDefault("CoinMarketCapKey", "")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error: %s", err))
	}

	return Config{viper.GetString("CoinMarketCapKey")}
}
