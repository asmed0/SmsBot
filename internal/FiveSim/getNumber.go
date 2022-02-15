package FiveSim

import (
	"encoding/json"
	"fmt"
	"github.com/getsentry/raven-go"
	"io/ioutil"
	"net/http"
	"strconv"
)

func getNumber(data *FiveSimSession) {
	url := fmt.Sprintf("https://5sim.net/v1/user/buy/activation/%s/%s/%s",
		data.Country,
		data.Operator,
		data.Product)

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s",data.ApiKey))

	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		fmt.Println(err)
		return
	}
	jsonPtr := &GetNumberResponse{}

	json.Unmarshal(body, jsonPtr)
	if jsonPtr.Phone != "" {
		data.ID = strconv.Itoa(jsonPtr.ID)
		data.Phone = jsonPtr.Phone
		data.Operator = jsonPtr.Operator
		data.Product = jsonPtr.Product
		data.Price = strconv.Itoa(jsonPtr.Price)
		data.Status = jsonPtr.Status
		data.Expires = jsonPtr.Expires
		data.Sms = jsonPtr.Sms
		data.CreatedAt = jsonPtr.CreatedAt
		data.Country = jsonPtr.Country
	} else {
		data.Operator = "virtual27"
		getNumber(data)
	}
}
