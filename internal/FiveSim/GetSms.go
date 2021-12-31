package FiveSim

import (
	"encoding/json"
	"fmt"
	"github.com/getsentry/raven-go"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetSms(session *FiveSimLastSession) string {
	url := fmt.Sprintf("https://5sim.net/v1/user/check/%s",
		session.ID)

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s",session.ApiKey))

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
	if len(jsonPtr.Sms) > 0{
		return jsonPtr.Sms[0].Text
	} else if strings.ToUpper(jsonPtr.Status) == "TIMEOUT"{
		return "ProviderErr"
	}else{
		return "Err"
	}
}
