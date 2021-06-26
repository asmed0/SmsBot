package SmsCodesIO

type Session struct {
	ApiKey      string `json:"apikey"`
	Country     string `json:"country"`
	ServiceID   string `json:"service_id"`
	SerciceName string `json:"sercice_name"`

	Number     string `json:"number"`
	SecurityID string `json:"security_id"`
}

type GetNumberResponse struct {
	Status     string `json:"Status"`
	Error      string `json:"Error"`
	Iso        string `json:"ISO"`
	Service    string `json:"Service"`
	SecurityID string `json:"SecurityId"`
	Number     string `json:"Number"`
	Rate       string `json:"Rate"`
}


type GetSmsResponse struct {
	Status  string `json:"Status"`
	Error   string `json:"Error"`
	Number  string `json:"Number"`
	Sms     string `json:"SMS"`
	Balance string `json:"Balance"`
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
	IsDisposed bool `bson:"is_disposed"`
}