package DiscordBot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/getsentry/raven-go"
	"os"
	"smsbot/configs"
	"smsbot/internal/Database"
	"smsbot/internal/SmsCodesIO"
	"smsbot/internal/Topup"
	"smsbot/internal/tools"
	"strconv"
	"strings"
)

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	data := configs.DBotConfigs()

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	//Ignore all messages outside #commands channel
	if m.ChannelID != os.Getenv("commands_channel") {
		return
	}

	//command slice handling
	cmds := tools.SliceSlicer(data.Commands)
	command := strings.TrimLeft(strings.ToLower(m.Content), data.Prefix)
	ckey, cbool := tools.Find(cmds, strings.Fields(command)[0])

	//opening a dm channel
	directMessage, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		return
	}

	embedMsg := &discordgo.MessageEmbed{
		Title:  "Unknown command, use !help command for more information on available commands!",
		Fields: []*discordgo.MessageEmbedField{},
		Provider: &discordgo.MessageEmbedProvider{
			URL:  "https://cdn.discordapp.com/icons/806511362251030558/244ed44d2ab37a59e37bb775de0d8fcb.png?size=256",
			Name: "SlotTalk SMSBOT",
		},
		Color: 16776960, //yellow color
		Footer: &discordgo.MessageEmbedFooter{
			Text:    "SmsBot by SlotTalk | Support? Sithed#4918",
			IconURL: "https://cdn.discordapp.com/icons/806511362251030558/244ed44d2ab37a59e37bb775de0d8fcb.png?size=256",
		},
	}

	if cbool == true { //fast way to determine if we want to handle this message or not
		switch ckey {
		case 0: //food command
			if !(Database.GetBalance(m.Author.ID) <= 0) {
				lastSession := Database.GetLastSession(m.Author.ID)
				if lastSession.Apikey != "" {
					isLastSessionDisposed := !lastSession.IsDisposed
					if isLastSessionDisposed {
						embedMsg.Title = "Please use the !code command to dispose previous number before requesting a new one"
						embedMsg.Color = 15158332 //red color
						go s.ChannelMessageSendEmbed(directMessage.ID, embedMsg)
						return
					}
				}

				number := Database.UpdateSession(m.Author.ID, SmsCodesIO.Init(), false)
				if number != "zerobal" {
					embedMsg.Title = "+" + number
					embedMsg.Description = "Use !code command to retrieve verification code\n *1 Token has been deducted from your balance"

					embedMsg.Color = 1752220 //aqua color
					go Database.UpdateBalance(-1, m.Author.ID)
					go s.ChannelMessageSendEmbed(directMessage.ID, embedMsg)
					go s.ChannelMessageSend(directMessage.ID, embedMsg.Title[3:len(embedMsg.Title)]) //stripping off +44
					return
				} else {
					embedMsg.Title = "No tokens left :("
					embedMsg.Description = "Use !topup command to purchase tokens!"
					embedMsg.Color = 15158332 //red color
					go s.ChannelMessageSendEmbed(directMessage.ID, embedMsg)
				}
			} else {
				embedMsg.Title = "You have no tokens :( Please use the !topup command first"
				embedMsg.Color = 15158332 //red color
				go s.ChannelMessageSendEmbed(directMessage.ID, embedMsg)
			}
		case 1: //code command
			returnedCode := SmsCodesIO.GetSms(Database.GetLastSession(m.Author.ID))
			if returnedCode == "Err" {
				embedMsg.Title = "Message not received yet, try again in a moment"
				embedMsg.Color = 15158332 //red color

				go s.ChannelMessageSendEmbed(directMessage.ID, embedMsg)
			} else {
				embedMsg.Title = returnedCode
				embedMsg.Description = "Use !balance command to check your balance!"
				embedMsg.Color = 3066993 //green color
				lastSession := Database.GetLastSession(m.Author.ID)
				go Database.UpdateSession(m.Author.ID, &SmsCodesIO.Session{
					ApiKey:      lastSession.Apikey,
					Country:     lastSession.Country,
					ServiceID:   lastSession.ServiceID,
					SerciceName: lastSession.ServiceName,
					Number:      lastSession.Number,
					SecurityID:  lastSession.SecurityID,
				}, true)

				go s.ChannelMessageSendEmbed(directMessage.ID, embedMsg)
			}
		case 2: //balance command
			embedMsg.Title = strconv.Itoa(Database.GetBalance(m.Author.ID)) + " Tokens left"
			embedMsg.Description = "Use !topup command to purchase more tokens!\n \n1 successfully retrieved verification code = 1 token redeemed!"
			embedMsg.Color = 10181046 //purple color

			go s.ChannelMessageSendEmbed(directMessage.ID, embedMsg)
		case 3: //topup command
			if len(strings.Fields(command)) > 1 {
				qty, _ := strconv.Atoi(strings.Fields(command)[1])
				if qty < 1 || qty > 50000 {
					qty = 10 //default to 10
				}
				embedMsg.URL = "https://checkout.stripe.com/pay/" + Topup.CreateCheckoutSession(m.Author.ID, qty)
				embedMsg.Title = "Click here to checkout " + strconv.Itoa(qty) + " tokens"
			} else {
				embedMsg.URL = "https://checkout.stripe.com/pay/" + Topup.CreateCheckoutSession(m.Author.ID, 10) //default 10 tokens
				embedMsg.Title = "Click here to checkout 10 tokens"

			}
			embedMsg.Description = "Tokens will automatically be added to your balance after!"
			embedMsg.Color = 15277667 //pink color
			go s.ChannelMessageSendEmbed(directMessage.ID, embedMsg)

		case 4: //fhelp command
			embedMsg.Title = "Available commands below"
			for i := 0; i < len(data.Commands); i++ {
				embedMsg.Fields = append(embedMsg.Fields, &discordgo.MessageEmbedField{
					Name:   data.Prefix + data.Commands[i][0],
					Value:  data.Commands[i][1],
					Inline: false,
				})
			} //looping existing commands

			go s.ChannelMessageSendEmbed(directMessage.ID, embedMsg)
		default:
			go s.ChannelMessageSendEmbed(m.ChannelID, embedMsg) //actually not needed since we trim messages but justtt incase
		}
	} else if m.ChannelID == directMessage.ID {
		go s.ChannelMessageSendEmbed(m.ChannelID, embedMsg) //if command is in dms person is obv talking to the bot
	}
}
