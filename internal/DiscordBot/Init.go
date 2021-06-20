package DiscordBot

func Init() {
	data := &DiscordData{
		Token:    "ODM0ODIwMDg3MDg1MDA2ODkw.YIGcyg.dmEPblcFyqOKT06Sog4Ary8zDGM",
		Prefix:   "!",
		Commands: []string{"food", "balance", "topup", "help"},
	}
	getConfig(true, data)
	start(data)
}
