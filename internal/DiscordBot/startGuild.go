package DiscordBot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/getsentry/raven-go"
	"os"
	"os/signal"
	"syscall"
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

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("DiscordBot is now monitoring #commands channel")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
