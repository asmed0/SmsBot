package configs

import (
	"os"
	"smsbot/internal/tools"
)

func SmsCodesIOConfigs(service string) *SmsCodesIOStruct {
	apikey := os.Getenv("smscodes_apikey")
	services := [][]string{
		//servicename, serviceid, country, price
		{"Foodora","462f7a96-98e9-44a5-9407-47d3104519bd", "UK"},
		{"Wolt","fe8f0b2c-510c-4d72-bbc0-327c141b4054", "UK"},
		{"Bolt","f8eb0240-dfb8-422b-8b9a-04f0c6fe6dee", "UK"},
		{"Tier","6d91fec7-24d5-4feb-98b9-1bebda232213", "UK"},
		{"Uber-Eats","c51f6b3b-7f87-454c-92bb-4874b02a3a7a", "SE"},
		{"Other","cf49a161-6626-4f3c-8e07-ebfffc9a0bab", "UK"},
	}
	//service slice handling
	srvc := tools.SliceSlicer(services)
	key, _ := tools.Find(srvc, service)


	config := &SmsCodesIOStruct{
		Apikey:      apikey,
		Country:     services[key][2],
		ServiceID:   services[key][1],
		ServiceName: services[key][0],
	}
	return config
}
