package DiscordBot

import (
	"github.com/bwmarrin/discordgo"
	"os"
	"smsbot/internal/Database"
	"smsbot/internal/FiveSim"
	"strconv"
)

func sendLogs(user string, embedMsg *discordgo.MessageEmbed, s *discordgo.Session, discordID string) {
	embed := *embedMsg//we dont want our logs embed to mess with userembed so new mem addr

	//opening a log channel
	logChannel := os.Getenv("log_channel")
	embed.Title = user + ":\n" + embed.Title

	lastSession := Database.GetLastSession(discordID)

	embed.Fields = []*discordgo.MessageEmbedField{
		{
			Name:   "User",
			Value:  "<@" + discordID + ">",
			Inline: true,
				},
		{
			Name:   "Balance",
			Value:  strconv.Itoa(Database.GetBalance(discordID)),
			Inline: true,
				},
		{
			Name:   "Service",
			Value:  lastSession.Product,
			Inline: true,
				},
		{
			Name:   "Number",
			Value:  lastSession.Phone,
			Inline: true,
				},
		{
			Name:   "Disposed",
			Value:  strconv.FormatBool(lastSession.IsDisposed),
			Inline: true,
				},
		{
			Name:   "SMS",
			Value:  FiveSim.GetSms(lastSession),
			Inline: true,
				},
	}
	embed.Footer.Text = "SmsBot by SlotTalk | Logs"
	s.ChannelMessageSendEmbed(logChannel, &embed)
}
