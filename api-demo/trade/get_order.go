package trade

import (
	"api-demo/util"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func GetOrder() {

	// set the user's apiKey/apiSecret
	apiKey := ""
	apiSecret := ""

	// set the domain url for calling OpenApi
	reqUrl := "https://openapi.xxx.com/sapi/v1/order"
	method := "GET"
	requestPath := "/sapi/v1/order"

	timestampMillisStr := strconv.FormatInt(util.UnixMillis(time.Now()), 10)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client {Transport:tr}
	req, err := http.NewRequest(method, reqUrl, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// set the parameters of the query order interface
	query := req.URL.Query()
	query.Add("orderId", "100001")
	//query.Add("newClientOrderId", "")
	query.Add("symbol", "LTCUSDT")
	queryString := query.Encode()

	// calculate the signature
	sign := util.GetSign(timestampMillisStr, method, requestPath, queryString, "", apiSecret)

	req.Header.Add("X-CH-SIGN", sign)
	req.Header.Add("X-CH-APIKEY", apiKey)
	req.Header.Add("X-CH-TS", timestampMillisStr)

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
	fmt.Println(fmt.Sprintf("GetOrder result: %s", string(body)))
}
