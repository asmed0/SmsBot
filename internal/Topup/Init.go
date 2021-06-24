package Topup

import (
	"github.com/stripe/stripe-go"
	"log"
	"net/http"
	"os"
)


func Init() {
	stripe.Key = os.Getenv("stripe_key")

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", http.FileServer(http.Dir(wd+"/internal/Topup/public")))
	http.HandleFunc("/success", checkSuccess) //handle success, add balance
	addr := ":" + os.Getenv("PORT")
	http.ListenAndServe(addr, nil)
}
