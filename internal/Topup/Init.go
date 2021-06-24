package Topup

import (
	"fmt"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"github.com/stripe/stripe-go/customer"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

func Init() {
	publishableKey := os.Getenv("PUBLISHABLE_KEY")
	stripe.Key = os.Getenv("SECRET_KEY")

	tmpls, _ := template.ParseFiles(filepath.Join("templates", "index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := tmpls.Lookup("index.html")
		tmpl.Execute(w, map[string]string{"Key": publishableKey})
	})

	http.HandleFunc("/charge", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		email := r.Form.Get("stripeEmail")
		customerParams := &stripe.CustomerParams{Email:&email}
		customerParams.SetSource(r.Form.Get("stripeToken"))

		newCustomer, err := customer.New(customerParams)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var amount int64 = 500
		currency := "usd"
		description := "sample charge"

		chargeParams := &stripe.ChargeParams{
			Amount:   &amount,
			Currency: &currency,
			Description:     &description,
			Customer: &newCustomer.ID,
		}

		if _, err := charge.New(chargeParams); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Charge completed successfully!")
	})

	http.ListenAndServe(":4567", nil)
}
