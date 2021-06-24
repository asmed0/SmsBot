package DiscordBot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"smsbot/internal/Database"
	"smsbot/internal/SmsCodesIO"
	"smsbot/internal/tools"
	"strconv"
	"strings"
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

	//filtering messages, shouldve used .trim() but idk bruh
	cmds := tools.SliceSlicer(data.Commands)
	key, bool := tools.Find(cmds, strings.ToLower(m.Content[1:len(m.Content)]))

	//opening a dm channel
	directMessage, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		log.Panic(err)
	}

	embedMsg := &discordgo.MessageEmbed{
		Title:       "Unknown command, use !help command for more information on available commands!",
		Description: "SmsBot by SlotTalk",
		Provider: &discordgo.MessageEmbedProvider{
			URL:  "https://cdn.discordapp.com/icons/806511362251030558/244ed44d2ab37a59e37bb775de0d8fcb.png?size=256",
			Name: "SlotTalk SMSBOT",
		},
		Color: 16776960, //yellow color
		Footer: &discordgo.MessageEmbedFooter{
			Text:    "Support? Ahmed#6969 or Sithed#4918",
			IconURL: "https://cdn.discordapp.com/icons/806511362251030558/244ed44d2ab37a59e37bb775de0d8fcb.png?size=256",
		},
	}

	if bool == true { //slicing a multidimensional slice has us jumping commands by +2 - ik there's better ways xdxd
		switch key {
		case 0: //food command
			embedMsg.Title = "+" + Database.UpdateSession(m.Author.ID, SmsCodesIO.Init())
			embedMsg.Description = "SmsBot by SlotTalk - Use !code command to retrieve verification code"

			embedMsg.Color = 1752220 //aqua color
			s.ChannelMessageSendEmbed(directMessage.ID, embedMsg)
			s.ChannelMessageSend(directMessage.ID, embedMsg.Title[3:len(embedMsg.Title)]) //stripping off +44

		case 2: //code command
			returnedCode := SmsCodesIO.GetSms(Database.GetLastSession(m.Author.ID))
			if strings.Contains(returnedCode, "not") {
				embedMsg.Title = returnedCode
				embedMsg.Description = "SmsBot by SlotTalk - Try again in a moment or resend the code!\n *Your balance is untouched"
				embedMsg.Color = 15158332 //red color

				s.ChannelMessageSendEmbed(directMessage.ID, embedMsg)
			} else {
				embedMsg.Title = returnedCode
				embedMsg.Description = "SmsBot by SlotTalk - Use !balance command to check your balance!\n *1 Token has been deducted from your balance"
				embedMsg.Color = 3066993 //green color

				go Database.UpdateBalance(-1, m.Author.ID)
				s.ChannelMessageSendEmbed(directMessage.ID, embedMsg)
			}
		case 4:
			embedMsg.Title = strconv.Itoa(Database.GetBalance(m.Author.ID)) + " Tokens left"
			embedMsg.Description = "SmsBot by SlotTalk - Use !topup command to purchase more tokens!\n \n1 successfully retrieved verification code = 1 token redeemed!"
			embedMsg.Color = 10181046 //purple color

			s.ChannelMessageSendEmbed(directMessage.ID, embedMsg)
		case 6:
			fmt.Println(data.Commands[key])
		case 8:
			fmt.Println(data.Commands[key])
		default:
			s.ChannelMessageSendEmbed(m.ChannelID, embedMsg) //actually not needed since we trim messages but justtt incase
		}
	} else {
		fmt.Println(cmds)
	}
}
