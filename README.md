# How to use
1. Replace `configreade/config.yaml`'s sample-key with real key
2. Compile
3. Run

# API
## `/api/v1/cryptocurrency/`
本API是用來查詢Bitcoin的現值。有兩個參數可以使用，分別是`currency`和`provider`。這兩個參數都是透過GET的方法所取得。
- `provider`: 選擇所要使用的provider，可使用的選項有`CoinMarketCap`、`CryptoCompare`、`CoinGecko`。若沒輸入預設為`CoinGecko`。
- `currency`: 選擇要查詢的貨幣單位，可使用的選項有`usd`、`twd`。若沒輸入預設為`twd`。

# Config Reader
可以透過config檔去改變程式的設定值，改完後要動啟程式才會生效。

以下為Config所代表的意義:
- `CoinMarketCapKey`: CoinMarkerCap API所需的key
- `CryptoCompareKey`: CryptoCompare API所需的key
- `SecondPerToken`: 對API provider所實做的流量控制系統為token bucket。若此參數值為N時，則每N秒加入一個token。
- `MaxSizeOfBucket`: 對API provider所實做的流量控制系統為token bucket。若此參數值為N時，則bucket的容量為N。
- `UserMaxQueryPerDay`: 對user的流量控制的參數。若此參數值為N時，則在24小時內user最多只能做出N次query。

# TODO
- Add database to save user query records
- Support history price query
- API provider side flow control add a cache to prevent frequently querry