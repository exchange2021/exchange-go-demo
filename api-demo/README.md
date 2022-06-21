##GO API Demo

####Getting started
`go run api_demo.go`

####Settings
configure these variables value in the go file will be executed;
like your apiKey/apiSecret/symbol...
```
func GetTicker() {
    query := req.URL.Query()
    query.Add("symbol", "LTCUSDT")
}
```

####Examples
#####Ticker
```
func main() {
    market.GetTicker()
}
```