package Topup

import (
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
		log.Printf("session.New: %v", err)
	}

	Database.AddPaymentSecret(discordID, session.ID)

	return session.ID + "#fidkdWxOYHwnPyd1blppbHNgWl91f2dMTUlQYVJsZGRMcTJwf2dLc29JdCcpJ2hsYXYnP34nYnBsYSc%2FJ2YwZDEzYGBmKGcyNzIoMTxkZig9Y2Y3KGc0YD1gZ2A0ZGRnN2QwZjxmZCcpJ2hwbGEnPydhMWM8Y2RkYShmMjBgKDFgNjMoPDNjMihkMTBmYDExNjJkNWBjYGdkMWYnKSd2bGEnPydjPD1gMTYxYChhMzZhKDFkMDMoPTA1Zig0NjM0NzY2ZjwyM2ZjMGQzPDYneCknZ2BxZHYnP15YKSdpZHxqcHFRfHVgJz8ndmxrYmlgWmxxYGgnKSd3YGNgd3dgd0p3bGJsayc%2FJ21xcXU%2FKippamZkaW1qdnE%2FNjU0NzYneCUl"
}
