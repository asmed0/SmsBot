package configs

import "os"

var token string
var cmds [][]string
func DBotConfigs() *DBotStruct{
	config := &DBotStruct{
		Token:    os.Getenv("discord_token"),
		AppID: os.Getenv("discord_appid"),
		Prefix:  "!",
		Commands: [][]string{ //{"command", `description`}
			{"food", `Use the !food command to get a Foodora number.`},
			{"balance", `Use the !balance command to check your SmsBot balance`},
			{"topup", `Use the !topup command to receive a payment link to purchase more SmsBot balance`},
			{"fhelp", `Use the !help command for more information on the available commands`}},
	}
	return config
}
