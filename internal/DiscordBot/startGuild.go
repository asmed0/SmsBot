package DiscordBot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/getsentry/raven-go"
	"log"
	"os"
	"os/signal"
)

func startGuild(data *DiscordData) {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + data.Token)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	dg.Identify.Presence = discordgo.GatewayStatusUpdate{
		Game:   discordgo.Activity{
			Name:          "Use the !fhelp command for more info on available commands!\nOpen a ticket for further support!",
		},
	}
	dg.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("Bot session is online!")
	})

	err = dg.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}
	defer dg.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Graceful shutdown")
}
