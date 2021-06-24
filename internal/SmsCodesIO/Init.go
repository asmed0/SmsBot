package SmsCodesIO

import "smsbot/configs"

func Init() *Session {
	config := configs.SmsCodesIOConfigs()
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
