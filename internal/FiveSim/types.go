package FiveSim

type FiveSimSession struct {
	ApiKey      string `json:"apikey"`
	ID     string `json:"id"`
	Phone      string `json:"phone"`
	Operator        string `json:"operator"`
	Product    string `json:"product"`
	Price string `json:"price"`
	Status     string `json:"status"`
	Expires       string `json:"expires"`
	Sms        []SmsSlice`json:"sms"`
	CreatedAt       string `json:"created_at"`
	Country       string `json:"country"`
}
type FiveSimLastSession struct {
	ApiKey      string `json:"apikey"`
	ID     string `json:"id"`
	Phone      string `json:"phone"`
	Operator        string `json:"operator"`
	Product    string `json:"product"`
	Price string `json:"price"`
	Status     string `json:"status"`
	Expires       string `json:"expires"`
	Sms        []SmsSlice`json:"sms"`
	CreatedAt       string `json:"created_at"`
	Country       string `json:"country"`
	IsDisposed bool `bson:"is_disposed"`
}

type GetNumberResponse struct {
	ID     int `json:"id"`
	Phone      string `json:"phone"`
	Operator        string `json:"operator"`
	Product    string `json:"product"`
	Price int `json:"price"`
	Status     string `json:"status"`
	Expires       string `json:"expires"`
	Sms        []SmsSlice`json:"sms"`
	CreatedAt       string `json:"created_at"`
	Country       string `json:"country"`
}


type GetSmsResponse struct {
	ID     int `json:"id"`
	Phone      string `json:"phone"`
	Operator        string `json:"operator"`
	Product    string `json:"product"`
	Price int `json:"price"`
	Status     string `json:"status"`
	Expires       string `json:"expires"`
	Sms        []SmsSlice`json:"sms"`
	CreatedAt       string `json:"created_at"`
	Country       string `json:"country"`
}

type SmsSlice struct{
	CreatedAt       string `json:"created_at"`
	Date       string `json:"date"`
	Sender       string `json:"sender"`
	Text       string `json:"text"`
	Code       string `json:"code"`

}

type UserData struct {
	ID          string          `bson:"_id"`
	DiscordID   string      `bson:"discord_id"`
	DiscordName string      `bson:"discord_name"`
	Balance     int         `bson:"balance"`
	LastSession FiveSimLastSession `bson:"last_session"`
}