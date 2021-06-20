package SmsCodesIO

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"syscall"
)

func getNumber(data *Session) {

	url := "https://admin.smscodes.io/api/sms/GetServiceNumber?key=" + data.ApiKey +
		"&iso=" + data.Country +
		"&serv=" + data.ServiceID

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
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
	jsonPtr := &GetNumberResponse{}

	json.Unmarshal(body, jsonPtr)
	if jsonPtr.Error == "Non" {
		data.Number = jsonPtr.Number
		data.SecurityID = jsonPtr.SecurityID
	} else {
		fmt.Println(jsonPtr.Error)
		syscall.Exit(-1)
	}
}
