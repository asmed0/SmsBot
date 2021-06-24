package main

import (
	"smsbot/internal/Database"
	"smsbot/internal/DiscordBot"
	"smsbot/internal/Topup"
)

func main() {

	Database.Init()
	go DiscordBot.Init()

	Topup.Init()

}
