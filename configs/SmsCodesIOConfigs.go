package configs

import "os"

func SmsCodesIOConfigs() *SmsCodesIOStruct{
	config := &SmsCodesIOStruct{
		Apikey:      os.Getenv("smscodes_apikey"),
		Country:     os.Getenv("smscodes_country"),
		ServiceID:   os.Getenv("smscodes_serviceid"),
		ServiceName: os.Getenv("smscodes_servicename"),
	}
	return config
}
