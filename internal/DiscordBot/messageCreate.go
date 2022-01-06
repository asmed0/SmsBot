package DiscordBot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/getsentry/raven-go"
	"log"
	"os"
	"smsbot/configs"
	"smsbot/internal/Database"
	"smsbot/internal/Topup"
	"smsbot/internal/tools"
	"strconv"
	"strings"
	"time"
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

	msg := &discordgo.MessageSend{
		Embeds:     nil,
		Components: []discordgo.MessageComponent{},
	}

	msg.Embeds = append(msg.Embeds, embedMsg)
	msg.Components = append(msg.Components, firstBtn)

	//Stripping command off prefix
	command := strings.TrimLeft(strings.ToLower(m.Content), data.Prefix)

	//Ignore all messages outside #commands channel if its not an admin
	commandsChan := os.Getenv("commands_channel")

	//if is Admin
	if m.ChannelID != directMessage.ID {
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

	if len(strings.Fields(command)) < 1 {
		return
	}

	switch strings.Fields(command)[0] {
	case "food": //food command
		embedMsg = embedCleaner() // cleaning the embed since previous usage
		embedMsg.Title = "We urge you to not request a number before you are ready to use it!"
		embedMsg.Description = "Click the green button below once you are ready :)"
		embedMsg.Color = 16776960
		go s.ChannelMessageSendComplex(directMessage.ID, msg)
	case "balance": //balance command
		embedMsg = embedCleaner() // cleaning the embed since previous usage
		embedMsg.Title = strconv.Itoa(Database.GetBalance(m.Author.ID)) + " Tokens left"
		embedMsg.Description = "Use !topup command to purchase more tokens!\n \n1 successfully retrieved verification code = 1 token redeemed!"
		embedMsg.Color = 10181046 //purple color

		go s.ChannelMessageSendEmbed(directMessage.ID, embedMsg)
	case "topup": //topup command
		embedMsg = embedCleaner() // cleaning the embed since previous usage
		if len(strings.Fields(command)) > 1 {
			qty, _ := strconv.Atoi(strings.Fields(command)[1])
			if qty < 5 || qty > 50000 { //no less than 5, no more than 50k
				qty = 5 //default to 5
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
		embedMsg = embedCleaner() // cleaning the embed since previous usage
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
		embedMsg = embedCleaner() // cleaning the embed since previous usage
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
				{
					Name:   "Previous balance",
					Value:  strconv.Itoa(prevBal),
					Inline: true,
				},
				{
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
			log.Println(fmt.Sprintf("[%v]unknown command: %s ", time.Now(), m.Content))//if member writes unknown message give them fhelp embed
		}
	}
}
