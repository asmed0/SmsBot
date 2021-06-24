package SmsCodesIO

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"syscall"
)

func GetSms(session *LastSession) string {
	url := "https://admin.smscodes.io/api/sms/GetSMSCode?key=" + session.Apikey +
		"&sid=" + session.SecurityID +
		"&number=" + session.Number

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return "Problem whilst retrieving sms code, contact support!"
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "Problem whilst retrieving sms code, contact support!"
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "Problem whilst retrieving sms code, contact support!"
	}

	jsonPtr := &GetSmsResponse{}

	json.Unmarshal(body, jsonPtr)
	if jsonPtr.Error == "Non" {
		return jsonPtr.Sms
	} else {
		fmt.Println(jsonPtr.Error)
		syscall.Exit(-1)
	}
	return "Problem whilst retrieving sms code, contact support!" //Unreachable code
}
