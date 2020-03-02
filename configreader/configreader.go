package configreader

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

// Config is a structure store config value
type Config struct {
	CoinMarketCapKey string
	CryptoCompareKey string
	SecondPerToken   uint
	MaxSizeOfBucket  uint
}

// ReadConfig will read config from current directory
func ReadConfig() Config {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	viper.SetDefault("CoinMarketCapKey", "")
	viper.SetDefault("CryptoCompareKey", "")
	viper.SetDefault("SecondPerToken", 0)
	viper.SetDefault("MaxSizeOfBucket", 0)
	viper.AddConfigPath(basepath)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error: %s", err))
	}

	return Config{
		CoinMarketCapKey: viper.GetString("CoinMarketCapKey"),
		CryptoCompareKey: viper.GetString("CryptoCompareKey"),
		SecondPerToken:   viper.GetUint("SecondPerToken"),
		MaxSizeOfBucket:  viper.GetUint("MaxSizeOfBucket"),
	}
}
