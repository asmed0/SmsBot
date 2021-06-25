package main

import (
	"github.com/getsentry/raven-go"
	"github.com/joho/godotenv"
	"os"
	"smsbot/internal/Database"
	"smsbot/internal/DiscordBot"
	"smsbot/internal/Topup"
)

func main() {
	godotenv.Load()

	raven.SetDSN(os.Getenv("sentry_url"))

	Database.Init()
	go DiscordBot.Init()
	Topup.Init()
}
