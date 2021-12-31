package configs

type DBStruct struct {
	Uri string
	User string
	Pass string
	Database string
	Collection string
}

type DBotStruct struct {
	Token string
	AppID string
	Prefix string
	Commands [][]string
}

type SmsCodesIOStruct struct {
	Apikey string
	Country string
	ServiceID string
	ServiceName string
	Price int
}

type FiveSimStruct struct {
	Apikey string
	Operator string
	Product string
	Price string
	Status string
	Expires string
	Sms []string
	CreatedAt string
	Country string
}
