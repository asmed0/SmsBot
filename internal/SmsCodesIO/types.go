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
