package Topup

import (
	"github.com/getsentry/raven-go"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/checkout/session"
	"log"
	"os"
	"smsbot/internal/Database"
	"strconv"
)

var qty int64
func CreateCheckoutSession(discordID string, qty int) string {

	item := &stripe.CheckoutSessionLineItemParams{
		Amount:   stripe.Int64(500),
		Currency: stripe.String(string(stripe.CurrencySEK)),
		Name:     stripe.String(strconv.Itoa(qty) + " SmsBot Tokens"),
		Quantity: stripe.Int64(int64(qty)),
		TaxRates: nil,
	}

	domain := os.Getenv("HOSTURL")
	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),

		LineItems:  []*stripe.CheckoutSessionLineItemParams{item},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(domain + "/success?session_id={CHECKOUT_SESSION_ID}&discord_id=" + discordID +
			"&qty=" + strconv.Itoa(qty)),
		CancelURL:  stripe.String(domain + "/cancel.html"),
	}

	session, err := session.New(params)

	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		log.Printf("session.New: %v", err)
	}

	Database.AddPaymentSecret(discordID, session.ID)

	return session.ID + "#" + os.Getenv("stripe_pk_encrypted")
}
