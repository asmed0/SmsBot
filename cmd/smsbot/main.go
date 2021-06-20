package main

import (
	"smsbot/internal/Database"
	"smsbot/internal/DiscordBot"
)

func main() {
	Database.Init()
	DiscordBot.Init()
}
