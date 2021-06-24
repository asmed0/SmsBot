package DiscordBot

func Init() {
	data := &DiscordData{
		Token:    "ODM0ODIwMDg3MDg1MDA2ODkw.YIGcyg.dmEPblcFyqOKT06Sog4Ary8zDGM",
		Prefix:   "!",
		Commands: [][]string{ //{"command", `description`}
			{"food", `Use the !food command to get a number. After requesting your sms code from Foodora please use the !code command.`},
			{"code", `Use the !code command to retrieve the latest received sms`},
			{"balance", `Use the !balance command to check your SmsBot balance`},
			{"topup", `Use the !topup command to receive a payment link to purchase more SmsBot balance`},
			{"help", `Use the !help command for more information on the available commands`}},
	}

	getConfig(true, data)
	start(data)
}
