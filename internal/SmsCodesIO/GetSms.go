package SmsCodesIO

import (
	"encoding/json"
	"fmt"
	"github.com/getsentry/raven-go"
	"io/ioutil"
	"net/http"
)

func GetSms(session *LastSession) string {
	url := "https://admin.smscodes.io/api/sms/GetSMSCode?key=" + session.Apikey +
		"&sid=" + session.SecurityID +
		"&number=" + session.Number

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		fmt.Println(err)
		return "Err"
	}
	res, err := client.Do(req)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		fmt.Println(err)
		return "Err"
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		fmt.Println(err)
		return "Err"
	}

	jsonPtr := &GetSmsResponse{}

	json.Unmarshal(body, jsonPtr)
	if jsonPtr.Error == "Non" && jsonPtr.Sms != "Message not received yet" {
		return jsonPtr.Sms
	} else {
		return "Err"
	}
}
