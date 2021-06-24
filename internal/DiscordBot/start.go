package DiscordBot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

func start(data *DiscordData) {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + data.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)
	//dg.AddHandler(newMember) //new members get greet message and bot instructions - will implement later

	dg.Identify.Intents = discordgo.IntentsGuildMessages
	dg.Identify.Intents = discordgo.IntentsDirectMessages //accept dms aswell


	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("DiscordBot is now running")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
