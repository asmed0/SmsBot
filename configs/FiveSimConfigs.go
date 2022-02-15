package configs

import (
	"os"
	"smsbot/internal/tools"
)

func FiveSimConfigs(service string) *FiveSimStruct {
	apikey := os.Getenv("5sim_apikey")
	services := [][]string{
		//servicename, opeator, country
		{"other","virtual15", "russia"},
	}
	//service slice handling
	srvc := tools.SliceSlicer(services)
	key, _ := tools.Find(srvc, service)


	config := &FiveSimStruct{
		Apikey:      apikey,
		Operator:   services[key][1],
		Product: services[key][0],
		Price: "",
		Status: "",
		Expires: "",
		Sms: nil,
		CreatedAt: "",
		Country:     services[key][2],
	}
	return config
}
