package DiscordBot

import (
	"github.com/bwmarrin/discordgo"
	"smsbot/internal/Database"
	"smsbot/internal/SmsCodesIO"
)

func getNumber(embedMsg *discordgo.MessageEmbed, discordID string, service string, price int) {
	if !(Database.GetBalance(discordID) <= 0) {
		lastSession := Database.GetLastSession(discordID)
		if lastSession.Apikey != "" {
			isLastSessionDisposed := !lastSession.IsDisposed
			if isLastSessionDisposed {
				embedMsg.Title = "Please use the !code command to dispose previous number before requesting a new one"
				embedMsg.Color = 15158332 //red color
				return
			}
		}

		number := Database.UpdateSession(discordID, SmsCodesIO.Init(service), false)
		if number != "zerobal" {
			embedMsg.Title = "+" + number
			embedMsg.Description = "Use !code command to retrieve verification code\n *1 Token has been deducted from your balance"

			embedMsg.Color = 1752220 //aqua color
			go Database.UpdateBalance(price, discordID)
			return
		} else {
			embedMsg.Title = "No tokens left :("
			embedMsg.Description = "Use !topup command to purchase tokens!"
			embedMsg.Color = 15158332 //red color
		}
	} else {
		embedMsg.Title = "You have no tokens :( Please use the !topup command first"
		embedMsg.Color = 15158332 //red color
	}
}
