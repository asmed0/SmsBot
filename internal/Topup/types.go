package Topup

type CheckoutData struct {
	Quantity  int64
	DiscordID string
}
type createCheckoutSessionResponse struct {
	SessionID string `json:"id"`
}

type UserData struct {
	ID          string          `bson:"_id"`
	DiscordID   string      `bson:"discord_id"`
	DiscordName string      `bson:"discord_name"`
	Balance     int         `bson:"balance"`
	LastSession LastSession `bson:"last_session"`
}

type LastSession struct {
	Apikey      string `bson:"apikey"`
	Country     string `bson:"country"`
	ServiceID   string `bson:"service_id"`
	ServiceName string `bson:"service_name"`
	Number      string `bson:"number"`
	SecurityID  string `bson:"security_id"`
}
