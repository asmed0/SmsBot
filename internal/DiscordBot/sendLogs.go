package DiscordBot

import (
	"github.com/bwmarrin/discordgo"
	"os"
	"smsbot/internal/Database"
	"smsbot/internal/SmsCodesIO"
	"strconv"
)

func sendLogs(user string, embedMsg *discordgo.MessageEmbed, s *discordgo.Session, discordID string) {
	embed := &discordgo.MessageEmbed{
		URL:         embedMsg.URL,
		Type:        embedMsg.Type,
		Title:       embedMsg.Title,
		Description: embedMsg.Description,
		Timestamp:   embedMsg.Timestamp,
		Color:       embedMsg.Color,
		Footer:      embedMsg.Footer,
		Image:       embedMsg.Image,
		Thumbnail:   embedMsg.Thumbnail,
		Video:       embedMsg.Video,
		Provider:    embedMsg.Provider,
		Author:      embedMsg.Author,
		Fields:      embedMsg.Fields,
	} //we dont want our logs embed to mess with userembed so new mem addr

	//opening a log channel
	logChannel := os.Getenv("log_channel")
	embed.Title = user + ":\n" + embed.Title

	lastSession := Database.GetLastSession(discordID)

	embed.Fields = []*discordgo.MessageEmbedField{
		&discordgo.MessageEmbedField{
			Name:   "User",
			Value:  "<@" + discordID + ">",
			Inline: true,
		},
		&discordgo.MessageEmbedField{
			Name:   "Balance",
			Value:  strconv.Itoa(Database.GetBalance(discordID)),
			Inline: true,
		},
		&discordgo.MessageEmbedField{
			Name:   "Service",
			Value:  lastSession.ServiceName,
			Inline: true,
		},
		&discordgo.MessageEmbedField{
			Name:   "Number",
			Value:  lastSession.Number,
			Inline: true,
		},
		&discordgo.MessageEmbedField{
			Name:   "Disposed",
			Value:  strconv.FormatBool(lastSession.IsDisposed),
			Inline: true,
		},
		&discordgo.MessageEmbedField{
			Name:   "SMS",
			Value:  SmsCodesIO.GetSms(lastSession),
			Inline: true,
		},
	}
	embed.Footer.Text = "SmsBot by SlotTalk | Logs"
	s.ChannelMessageSendEmbed(logChannel, embed)
}
