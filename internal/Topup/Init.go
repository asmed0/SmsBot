package Topup

import (
	"encoding/json"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/checkout/session"
	"log"
	"net/http"
	"os"
	"strconv"
)


func Init() {
	stripe.Key = os.Getenv("stripe_key")

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/lol", http.FileServer(http.Dir(wd+"/internal/Topup/public")))
	http.HandleFunc("/success", checkSuccess) //handle success, add balance
	http.HandleFunc("/create-checkout-session", createCheckoutSession)
	addr := ":" + os.Getenv("PORT")
	http.ListenAndServe(addr, nil)
}

func createCheckoutSession(w http.ResponseWriter, req *http.Request) {
	item := &stripe.CheckoutSessionLineItemParams{
		Amount:   stripe.Int64(500),
		Currency: stripe.String(string(stripe.CurrencySEK)),
		Name:     stripe.String(strconv.Itoa(10) + " SmsBot Tokens"),
		Quantity: stripe.Int64(int64(10)),
		TaxRates: nil,
	}
	domain := "http://localhost:4242"
	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		LineItems: []*stripe.CheckoutSessionLineItemParams{item},
		Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(domain + "/success.html"),
		CancelURL: stripe.String(domain + "/cancel.html"),
	}
	session, err := session.New(params)
	if err != nil {
		log.Printf("session.New: %v", err)
	}
	data := createCheckoutSessionResponse{
		SessionID: session.ID,
	}
	js, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
