package SmsCodesIO

import (
	"smsbot/configs"
)

func Init(service string) *Session {
	config := configs.SmsCodesIOConfigs(service)
	session := &Session{
		ApiKey:      config.Apikey,
		Country:     config.Country,
		ServiceID:   config.ServiceID,
		SerciceName: config.ServiceName,
		Number:      "",
		SecurityID:  "",
	}
	getNumber(session)
	return session
}
