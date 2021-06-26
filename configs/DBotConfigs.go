package configs

import "os"

var token string
var cmds [][]string
func DBotConfigs() *DBotStruct{
	config := &DBotStruct{
		Token:    os.Getenv("discord_token"),
		Prefix:  "!",
		Commands: [][]string{ //{"command", `description`}
			{"food", `Use the !food command to get a Foodora number.`},
			{"wolt", `Use the !wolt command to get a Wolt number.`},
			{"bolt", `Use the !bolt command to get a Bolt number.`},
			{"tier", `Use the !Tier command to get a Tier number.`},
			{"code", `Use the !code command to retrieve your verification code`},
			{"balance", `Use the !balance command to check your SmsBot balance`},
			{"topup", `Use the !topup command to receive a payment link to purchase more SmsBot balance`},
			{"fhelp", `Use the !help command for more information on the available commands`}},
	}
	return config
}
