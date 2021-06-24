package DiscordBot

import "smsbot/configs"

func Init() {
	config := configs.DBotConfigs()
	data := &DiscordData{
		Token:    config.Token,
		Prefix:   config.Prefix,
		Commands: config.Commands,
	}
	start(data)
}
