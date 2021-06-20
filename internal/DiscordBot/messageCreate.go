package DiscordBot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"smsbot/internal/Database"
	"smsbot/internal/SmsCodesIO"
	"smsbot/internal/tools"
)

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	data := getConfig(false, &DiscordData{})

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	//filtering messages, couldve used .trim() lol
	key, bool := tools.Find(data.Commands, m.Content[1:len(m.Content)])

	if bool == true {
		switch key {
		case 0: //food command
			directMessage, err := s.UserChannelCreate(m.Author.ID)
			if err != nil {
				log.Panic(err)
			}

			s.ChannelMessageSend(directMessage.ID, Database.UpdateSession(m.Author.ID, SmsCodesIO.Init()))

		case 1:
			fmt.Println(data.Commands[key])
		case 2:
			fmt.Println(data.Commands[key])
		case 3:
			fmt.Println(data.Commands[key])
		case 4:
			fmt.Println(data.Commands[key])
		}
	}
}
