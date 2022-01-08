package tools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func DiscountGen() string {
	url := "https://apiv3.rabble.com/checkouts"
	method := "POST"

	payload := strings.NewReader(`{"offer_id":31768,"spot_id":6685,"location":{"lat":59.3293,"lng":18.0686}}`)

	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Host", "apiv3.rabble.com")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("sec-ch-ua", "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"96\", \"Google Chrome\";v=\"96\"")
	req.Header.Add("Content-type", "application/json")
	req.Header.Add("Application", "web; os=macos; version=8.4.4")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Origin", "https://www.rabble.se")
	req.Header.Add("Sec-Fetch-Site", "cross-site")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Referer", "https://www.rabble.se/")
	req.Header.Add("Accept-Language", "en-US,en;q=0.9,sv;q=0.8")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	jsonPtr := &RabbleResponse{}

	json.Unmarshal(body, jsonPtr)

	return jsonPtr.Checkout.Coupon.Code
}

type RabbleResponse struct {
	Checkout struct {
		UpdatedAt time.Time `json:"updated_at"`
		Coupon    struct {
			Visualization string    `json:"visualization"`
			CheckoutID    int       `json:"checkout_id"`
			Code          string    `json:"code"`
			UpdatedAt     time.Time `json:"updated_at"`
			CreatedAt     time.Time `json:"created_at"`
			ID            int       `json:"id"`
			CodeText      string    `json:"code_text"`
			RedemptionURI string    `json:"redemption_uri"`
			OfferID       int       `json:"offer_id"`
		} `json:"coupon"`
		CreatedAt time.Time `json:"created_at"`
		Location  struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"location"`
		ID          int  `json:"id"`
		Invalidated bool `json:"invalidated"`
	} `json:"checkout"`
}
