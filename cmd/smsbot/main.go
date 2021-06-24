package main

import (
	"github.com/joho/godotenv"
	"smsbot/internal/Database"
	"smsbot/internal/DiscordBot"
	"smsbot/internal/Topup"
)

func main() {
	godotenv.Load()
	Database.Init()
	go DiscordBot.Init()

	Topup.Init()
}
