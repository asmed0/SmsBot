package SmsCodesIO

import (
	"smsbot/configs"
)

func Init(service string) *SmsCodesSession {
	config := configs.SmsCodesIOConfigs(service)
	session := &SmsCodesSession{
		ApiKey:      config.Apikey,
		Country:     config.Country,
		ServiceID:   config.ServiceID,
		ServiceName: config.ServiceName,
		Number:      "",
		SecurityID:  "",
	}
	getNumber(session)
	return session
}
