package trade

import (
	"api-demo/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type OrderReqParams struct {
	Symbol string `json:"symbol"`
	Volume string `json:"volume"`
	Side string `json:"side"`
	Type string `json:"type"`
	Price string `json:"price"`
	NewClientOrderId string `json:"newClientOrderId"`
	RecvWindow string `json:"recvWindow"`
}

func NewOrder() {

	// set the user's apiKey/apiSecret
	apiKey := ""
	apiSecret := ""

	// set the domain url for calling OpenApi
	reqUrl := "https://openapi.xxx.com/sapi/v1/order"
	method := "POST"
	requestPath := "/sapi/v1/order"
	queryString := ""

	// set the parameters of the create order interface
	var OrderReqParams = &OrderReqParams{
		Symbol: "LTCUSDT",
		Volume: "10",
		Side: "BUY",
		Type: "LIMIT",
		Price: "60",
		//NewClientOrderId: "",
		//RecvWindow: "",
	}
	orderReqParamsBytes, _ := json.Marshal(OrderReqParams)
	requestBody := string(orderReqParamsBytes)

	timestampMillisStr := strconv.FormatInt(util.UnixMillis(time.Now()), 10)

	// calculate the signature
	sign := util.GetSign(timestampMillisStr, method, requestPath, queryString, requestBody, apiSecret)

	payload := strings.NewReader(requestBody)

	client := &http.Client {}
	req, err := http.NewRequest(method, reqUrl, payload)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("X-CH-SIGN", sign)
	req.Header.Add("X-CH-APIKEY", apiKey)
	req.Header.Add("X-CH-TS", timestampMillisStr)
	req.Header.Add("Content-Type", "application/json")

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
	fmt.Println(fmt.Sprintf("NewOrder result: %s", string(body)))
}
