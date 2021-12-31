package FiveSim

import (
	"smsbot/configs"
)

func Init(service string) *FiveSimSession {
	config := configs.FiveSimConfigs(service)
	session := &FiveSimSession{
		ApiKey:      config.Apikey,
		ID:  "",
		Phone:      "",
		Operator:   config.Operator,
		Product: config.Product,
		Price: config.Price,
		Status: config.Status,
		Expires: config.Expires,
		CreatedAt: config.CreatedAt,
		Country:     config.Country,
	}
	getNumber(session)
	return session
}
