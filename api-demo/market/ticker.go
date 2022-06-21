package market

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetTicker() {

	// set the domain url for calling OpenApi
	reqUrl := "https://openapi.xxx.com/sapi/v1/ticker"
	method := "GET"

	client := &http.Client {}
	req, err := http.NewRequest(method, reqUrl, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// set the parameters of the query market interface
	query := req.URL.Query()
	query.Add("symbol", "LTCUSDT")

	req.URL.RawQuery = query.Encode()

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	// print return data
	fmt.Println(fmt.Sprintf("GetTicker result: %s", string(body)))
}
