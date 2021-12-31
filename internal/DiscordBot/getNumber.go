package DiscordBot

import (
	"github.com/bwmarrin/discordgo"
	"smsbot/internal/Database"
	"smsbot/internal/FiveSim"
	"strconv"
)

func getNumber(embedMsg *discordgo.MessageEmbed, discordID string, service string, price int) {
	if !(Database.GetBalance(discordID) <= 0) {
		lastSession := Database.GetLastSession(discordID)
		if lastSession.ApiKey != "" {
			isLastSessionDisposed := !lastSession.IsDisposed
			if isLastSessionDisposed {
				embedMsg.Title = "Please use the !code command to dispose previous number before requesting a new one"
				embedMsg.Color = 15158332 //red color
				return
			}
		}

		number := Database.UpdateSession(discordID, FiveSim.Init(service), false)
		if number != "zerobal" {
			embedMsg.Title = number
			embedMsg.Description = "Use !code command to retrieve verification code\n *" + strconv.Itoa(price) + " Token(s) has been deducted from your balance"

			embedMsg.Color = 1752220 //aqua color
			go Database.UpdateBalance(price, discordID)
			return
		} else {
			embedMsg.Title = "No tokens left :("
			embedMsg.Description = "Use !topup command to purchase tokens!"
			embedMsg.Color = 15158332 //red color
		}
	} else {
		embedMsg.Title = "You have no tokens left :("
		embedMsg.Color = 15158332 //red color
	}
}
