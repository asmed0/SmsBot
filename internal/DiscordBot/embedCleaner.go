package DiscordBot

import "github.com/bwmarrin/discordgo"

func embedCleaner(embedPtr *discordgo.MessageEmbed){
	embedPtr = &discordgo.MessageEmbed{
		Title:  "Unknown command, use !fhelp command for more information on available commands!",
		Fields: []*discordgo.MessageEmbedField{},
		URL: "",
		Provider: &discordgo.MessageEmbedProvider{
			URL:  "https://cdn.discordapp.com/icons/806511362251030558/244ed44d2ab37a59e37bb775de0d8fcb.png?size=256",
			Name: "SlotTalk SMSBOT",
		},
		Color: 16776960, //yellow color
		Footer: &discordgo.MessageEmbedFooter{
			Text:    "SmsBot by SlotTalk | Support? Open a ticket!",
			IconURL: "https://cdn.discordapp.com/icons/806511362251030558/244ed44d2ab37a59e37bb775de0d8fcb.png?size=256",
		},
	}
}
