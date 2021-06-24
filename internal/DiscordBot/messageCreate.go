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

	//filtering messages, couldve used .trim() lol
	cmds := tools.SliceSlicer(data.Commands)
	key, bool := tools.Find(cmds, m.Content[1:len(m.Content)])

	//opening a dm channel
	directMessage, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		log.Panic(err)
	}

	if bool == true { //slicing a multidimensional slice has us jumping commands by +2 - ik there's better ways xdxd
		switch key {
		case 0: //food command
			embedMsg := &discordgo.MessageEmbed{
				Title:       "+" + Database.UpdateSession(m.Author.ID, SmsCodesIO.Init()),
				Description: "SmsBot by SlotTalk - Use !code command to retrieve verification code",
				Color:       1752220, //aqua color
				Footer: &discordgo.MessageEmbedFooter{
					Text:    "Support? Ahmed#6969 or Sithed#4918",
					IconURL: "https://cdn.discordapp.com/icons/806511362251030558/244ed44d2ab37a59e37bb775de0d8fcb.png?size=256",
				},
			}

			s.ChannelMessageSendEmbed(directMessage.ID, embedMsg)
			s.ChannelMessageSend(directMessage.ID, embedMsg.Title[3:len(embedMsg.Title)]) //stripping off +44

		case 2: //code command
		returnedCode := SmsCodesIO.GetSms(Database.GetLastSession(m.Author.ID))
		if strings.Contains(returnedCode, "not"){
			embedMsg := &discordgo.MessageEmbed{
				Title:       returnedCode,
				Description: "SmsBot by SlotTalk - Try again in a moment or resend the code!\n *Your balance is untouched",
				Color:       15158332, //red color
				Footer: &discordgo.MessageEmbedFooter{
					Text:    "Support? Ahmed#6969 or Sithed#4918",
					IconURL: "https://cdn.discordapp.com/icons/806511362251030558/244ed44d2ab37a59e37bb775de0d8fcb.png?size=256",
				},
			}
			s.ChannelMessageSendEmbed(directMessage.ID, embedMsg)
		}else{
			embedMsg := &discordgo.MessageEmbed{
				Title:       returnedCode,
				Description: "SmsBot by SlotTalk - Use !balance command to check your balance!\n *1 Token has been deducted from your balance",
				Color:       3066993, //green color
				Footer: &discordgo.MessageEmbedFooter{
					Text:    "Support? Ahmed#6969 or Sithed#4918",
					IconURL: "https://cdn.discordapp.com/icons/806511362251030558/244ed44d2ab37a59e37bb775de0d8fcb.png?size=256",
				},
			}
			go Database.UpdateBalance(-1,m.Author.ID )
			s.ChannelMessageSendEmbed(directMessage.ID, embedMsg)
		}
		case 4:
			embedMsg := &discordgo.MessageEmbed{
				Title:       strconv.Itoa(Database.GetBalance(m.Author.ID)) + " Tokens left",
				Description: "SmsBot by SlotTalk - Use !topup command to purchase more tokens!\n \n1 successfully retrieved verification code = 1 token redeemed!",
				Color:       10181046, //purple color
				Footer: &discordgo.MessageEmbedFooter{
					Text:    "Support? Ahmed#6969 or Sithed#4918",
					IconURL: "https://cdn.discordapp.com/icons/806511362251030558/244ed44d2ab37a59e37bb775de0d8fcb.png?size=256",
				},
			}
			s.ChannelMessageSendEmbed(directMessage.ID, embedMsg)
		case 6:
			fmt.Println(data.Commands[key])
		case 8:
			fmt.Println(data.Commands[key])
		}
	} else {
			fmt.Println(cmds)
	}
}
