package main

import (
	"github.com/joho/godotenv"
	"log"
	"smsbot/internal/Database"
	"smsbot/internal/DiscordBot"
	"smsbot/internal/Topup"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	Database.Init()
	go DiscordBot.Init()

	Topup.Init()
}
