package DiscordBot

import (
	"errors"
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

var isAdmin bool

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	raven.SetUserContext(&raven.User{ID: m.Author.ID, Username: m.Author.Username})
	data := configs.DBotConfigs()

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	//Ignore other bots messages
	if m.Author.Bot {
		return
	}


	//opening a dm channel
	directMessage, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil) //shouldnt happen but ait
		return
	}

	embedMsg := &discordgo.MessageEmbed{
		Title:  "Unknown command, use !fhelp command for more information on available commands!",
		Fields: []*discordgo.MessageEmbedField{},
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

	//Stripping command off prefix
	command := strings.TrimLeft(strings.ToLower(m.Content), data.Prefix)

	//Ignore all messages outside #commands channel if its not an admin
	commandsChan := os.Getenv("commands_channel")

	//if is Admin
	if m.ChannelID != directMessage.ID{
		_, isAdmin = tools.Find(m.Member.Roles, os.Getenv("admin_role"))
	}

	if m.ChannelID == directMessage.ID {
		if command != "code" {
			return
		}
	} else if !isAdmin {
		if m.ChannelID != commandsChan {
			return
		}
	}

	if len(strings.Fields(command)) < 1{
		raven.CaptureErrorAndWait(errors.New("No command error: " + strings.Fields(command)[0]), nil)
		return
	}

	switch strings.Fields(command)[0] {
	case "food": //food command
		getNumber(embedMsg, m.Author.ID, "Foodora", -1)
		go s.ChannelMessageSendEmbed(directMessage.ID, embedMsg)
		go sendLogs(m.Author.Username, embedMsg, s, m.Author.ID)
	case "wolt":
		getNumber(embedMsg, m.Author.ID, "Wolt", -2)
		go s.ChannelMessageSendEmbed(directMessage.ID, embedMsg)
		go sendLogs(m.Author.Username, embedMsg, s, m.Author.ID)
	case "bolt":
		getNumber(embedMsg, m.Author.ID, "Bolt", -2)
		go s.ChannelMessageSendEmbed(directMessage.ID, embedMsg)
		go sendLogs(m.Author.Username, embedMsg, s, m.Author.ID)
	case "tier":
		getNumber(embedMsg, m.Author.ID, "Tier", -2)
		go s.ChannelMessageSendEmbed(directMessage.ID, embedMsg)
		go sendLogs(m.Author.Username, embedMsg, s, m.Author.ID)
	case "code": //code command
		returnedCode := SmsCodesIO.GetSms(Database.GetLastSession(m.Author.ID))
		if returnedCode == "Err" {
			embedMsg.Title = "Message not received yet, try again in a moment"
			embedMsg.Color = 15158332 //red color

			go s.ChannelMessageSendEmbed(directMessage.ID, embedMsg)
			go sendLogs(m.Author.Username, embedMsg, s, m.Author.ID)
		} else if returnedCode == "ProviderErr" {
			embedMsg.Title = "Seems like our provider had an error, sorry for this"
			embedMsg.Description = "Your balance has been reimbursed"
			embedMsg.Color = 15158332 //red color
			go Database.UpdateBalance(2, m.Author.ID)
			go s.ChannelMessageSendEmbed(directMessage.ID, embedMsg)
			go sendLogs(m.Author.Username, embedMsg, s, m.Author.ID)
		} else {
			embedMsg.Title = returnedCode
			embedMsg.Description = "Use !balance command to check your balance!"
			embedMsg.Color = 3066993 //green color

			go s.ChannelMessageSendEmbed(directMessage.ID, embedMsg)
			go sendLogs(m.Author.Username, embedMsg, s, m.Author.ID)
		}

		lastSession := Database.GetLastSession(m.Author.ID)
		Database.UpdateSession(m.Author.ID, &SmsCodesIO.Session{
			ApiKey:      lastSession.Apikey,
			Country:     lastSession.Country,
			ServiceID:   lastSession.ServiceID,
			SerciceName: lastSession.ServiceName,
			Number:      lastSession.Number,
			SecurityID:  lastSession.SecurityID,
		}, true) //disposing our number

	case "balance": //balance command
		embedMsg.Title = strconv.Itoa(Database.GetBalance(m.Author.ID)) + " Tokens left"
		embedMsg.Description = "Use !topup command to purchase more tokens!\n \n1 successfully retrieved verification code = 1 token redeemed!"
		embedMsg.Color = 10181046 //purple color

		go s.ChannelMessageSendEmbed(directMessage.ID, embedMsg)
	case "topup": //topup command
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

	case "fhelp": //fhelp command
		embedMsg.Title = "Available commands below"
		for i := 0; i < len(data.Commands); i++ {
			embedMsg.Fields = append(embedMsg.Fields, &discordgo.MessageEmbedField{
				Name:   data.Prefix + data.Commands[i][0],
				Value:  data.Commands[i][1],
				Inline: true,
			})
		} //looping existing commands

		go s.ChannelMessageSendEmbed(directMessage.ID, embedMsg)

	case "addtokens": //this is an admin only command! //example: !addtokens 50 @Santa
		if isAdmin {
			if len(m.Mentions) < 1 { //no user specified err
				embedMsg.Title = "No user specified"
				embedMsg.Description = "!addtokens command example: !addtokens 50 @Santa"
				go s.ChannelMessageSendEmbed(directMessage.ID, embedMsg)
				return
			}
			prevBal := Database.GetBalance(m.Mentions[0].ID)
			toAdd, err := strconv.Atoi(strings.Fields(command)[1])
			if err != nil {
				embedMsg.Title = "Incorrect amount to add, try again.."
				embedMsg.Description = "!addtokens command example: !addtokens 50 @Santa"
				go s.ChannelMessageSendEmbed(directMessage.ID, embedMsg)
				return
			}

			Database.UpdateBalance(toAdd, m.Mentions[0].ID)
			embedMsg.Title = m.Mentions[0].Username + "'s token balance has been updated"
			embedMsg.Fields = []*discordgo.MessageEmbedField{
				&discordgo.MessageEmbedField{
					Name:   "Previous balance",
					Value:  strconv.Itoa(prevBal),
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name:   "New balance",
					Value:  strconv.Itoa(Database.GetBalance(m.Mentions[0].ID)),
					Inline: true,
				},
			}
			//opening a dm channel to user
			userDm, err := s.UserChannelCreate(m.Mentions[0].ID)
			if err != nil {
				raven.CaptureErrorAndWait(err, nil) //shouldnt happen but ait
				return
			}
			go s.ChannelMessageSendEmbed(userDm.ID, embedMsg)
			go s.ChannelMessageSendEmbed(os.Getenv("log_channel"), embedMsg)
		}

	default:

		if !isAdmin {
			go s.ChannelMessageSendEmbed(m.ChannelID, embedMsg) //if member writes unknown message give them fhelp embed
		}
	}
}
