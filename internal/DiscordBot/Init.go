package DiscordBot

import "smsbot/configs"

func Init() {
	config := configs.DBotConfigs()
	data := &DiscordData{
		Token:    config.Token,
		AppID: config.AppID,
		Prefix:   config.Prefix,
		Commands: config.Commands,
	}
	//go startDM(data) //extra goroutine for dms
	go startCommands(data) //starting cmds seperately
	startGuild(data)

}
