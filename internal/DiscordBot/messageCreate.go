package DiscordBot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/getsentry/raven-go"
	"os"
	"smsbot/configs"
	"smsbot/internal/Database"
	"smsbot/internal/FiveSim"
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

	firstBtn := discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{
			discordgo.Button{
				Emoji: discordgo.ComponentEmoji{
					Name: "🎢",
				},
				Label:    "I'm ready!",
				Style:    discordgo.SuccessButton,
				CustomID: "imready",
			},
		},
	}

	componentsHandlers := map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"imready": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			getNumber(embedMsg, m.Author.ID, "other", -1)
			if embedMsg.Color != 15158332 {
				err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "",
						Flags:   1 << 6,
						Embeds:  []*discordgo.MessageEmbed{embedMsg},
						Components: []discordgo.MessageComponent{
							discordgo.ActionsRow{
								Components: []discordgo.MessageComponent{
									discordgo.Button{
										Emoji: discordgo.ComponentEmoji{
											Name: "🎰",
										},
										Label:    "Request code!",
										Style:    discordgo.PrimaryButton,
										CustomID: "getcode",
									},
								},
							},
						},
					},
				})
				if err != nil {
					panic(err)
				}
			} else {
				embedMsg.Description = "If you wish to topup less than 10 tokens (default amount) you can use the !topup command"
				err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "",
						Flags:   1 << 6,
						Embeds:  []*discordgo.MessageEmbed{embedMsg},
						Components: []discordgo.MessageComponent{
							discordgo.ActionsRow{
								Components: []discordgo.MessageComponent{
									discordgo.Button{
										Emoji: discordgo.ComponentEmoji{
											Name: "🏦",
										},
										Label:    "Click here to topup!",
										Style:    discordgo.LinkButton,
										URL: "https://checkout.stripe.com/pay/" + Topup.CreateCheckoutSession(m.Author.ID, 10),
									},
								},
							},
						},
					},
				})
				if err != nil {
					panic(err)
				}
			}
		},
		"getcode": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			returnedCode := FiveSim.GetSms(Database.GetLastSession(m.Author.ID))
			if returnedCode == "Err" {
				embedMsg.Title = "Message not received yet, try again in a moment"
				embedMsg.Description = ""
				embedMsg.Color = 15158332 //red color
				go sendLogs(m.Author.Username, embedMsg, s, m.Author.ID)
				err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "",
						Flags:   1 << 6,
						Embeds:  []*discordgo.MessageEmbed{embedMsg},
						Components: []discordgo.MessageComponent{
							discordgo.ActionsRow{
								Components: []discordgo.MessageComponent{
									discordgo.Button{
										Emoji: discordgo.ComponentEmoji{
											Name: "🥠",
										},
										Label:    "Request code again!",
										Style:    discordgo.SecondaryButton,
										CustomID: "getcode",
									},
								},
							},
						},
					},
				})
				if err != nil {
					panic(err)
				}
			} else if returnedCode == "ProviderErr" {
				embedMsg.Title = "Seems like you took to long to use the number, or our provider had an error."
				embedMsg.Description = "Your balance has been reimbursed"
				embedMsg.Color = 15158332 //red color
				go Database.UpdateBalance(2, m.Author.ID)
				go sendLogs(m.Author.Username, embedMsg, s, m.Author.ID)
				err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "",
						Flags:   1 << 6,
						Embeds:  []*discordgo.MessageEmbed{embedMsg},
						Components: []discordgo.MessageComponent{
							discordgo.ActionsRow{
								Components: []discordgo.MessageComponent{
									discordgo.Button{
										Emoji: discordgo.ComponentEmoji{
											Name: "♻️",
										},
										Label:    "Request a new number!",
										Style:    discordgo.PrimaryButton,
										CustomID: "imready",
									},
								},
							},
						},
					},
				})
				if err != nil {
					panic(err)
				}
			} else {
				embedMsg.Title = returnedCode
				embedMsg.Description = "Thank you for using our services"
				embedMsg.Color = 3066993 //green color

				go sendLogs(m.Author.Username, embedMsg, s, m.Author.ID)
				err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "",
						Flags:   1 << 6,
						Embeds:  []*discordgo.MessageEmbed{embedMsg},
						Components: []discordgo.MessageComponent{
							discordgo.ActionsRow{
								Components: []discordgo.MessageComponent{
									discordgo.Button{
										Emoji: discordgo.ComponentEmoji{
											Name: "💰",
										},
										Label:    "Check balance!",
										Style:    discordgo.PrimaryButton,
										CustomID: "getbalance",
									},
									discordgo.Button{
										Emoji: discordgo.ComponentEmoji{
											Name: "♻️",
										},
										Label:    "Request new code!",
										Style:    discordgo.SecondaryButton,
										CustomID: "getcode",
									},
								},
							},
						},
					},
				})
				if err != nil {
					panic(err)
				}
			}
			lastSession := Database.GetLastSession(m.Author.ID)
			if returnedCode != "Err" && returnedCode != "ProverErr" {
				lastSession.Sms = append(lastSession.Sms, FiveSim.SmsSlice{Text: returnedCode})
			}
			Database.UpdateSession(m.Author.ID, &FiveSim.FiveSimSession{
				ApiKey:    lastSession.ApiKey,
				ID:        lastSession.ID,
				Phone:     lastSession.Phone,
				Operator:  lastSession.Country,
				Product:   lastSession.Product,
				Price:     lastSession.Price,
				Status:    lastSession.Status,
				Expires:   lastSession.Country,
				Sms:       lastSession.Sms,
				CreatedAt: lastSession.Country,
				Country:   lastSession.Country,
			}, true) //disposing our number
		},
		"getbalance": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			embedMsg.Title = strconv.Itoa(Database.GetBalance(m.Author.ID)) + " Tokens left"
			embedMsg.Description = "Use !topup command to purchase more tokens!\n \n1 successfully retrieved verification code = 1 token redeemed!"
			embedMsg.Color = 10181046 //purple color
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "",
					Flags:   1 << 6,
					Embeds:  []*discordgo.MessageEmbed{embedMsg},
				},
			})
			if err != nil {
				panic(err)
			}
		},
	}
	// Components are part of interactions, so we register InteractionCreate handler
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch i.Type {
		case discordgo.InteractionMessageComponent:
			if h, ok := componentsHandlers[i.MessageComponentData().CustomID]; ok {
				h(s, i)
			}
		}
	})
	_, err = s.ApplicationCommandCreate(data.AppID, "", &discordgo.ApplicationCommand{
		Name:        "food",
		Description: "Request a number",
	})

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
		embedMsg.Title = "We urge you to not request a number before you are ready to use it!"
		embedMsg.Description = "Click the green button below once you are ready :)"
		go s.ChannelMessageSendComplex(directMessage.ID, msg)
	case "balance": //balance command
		embedMsg.Title = strconv.Itoa(Database.GetBalance(m.Author.ID)) + " Tokens left"
		embedMsg.Description = "Use !topup command to purchase more tokens!\n \n1 successfully retrieved verification code = 1 token redeemed!"
		embedMsg.Color = 10181046 //purple color

		go s.ChannelMessageSendEmbed(directMessage.ID, embedMsg)
	case "topup": //topup command
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
			go s.ChannelMessageSendEmbed(m.ChannelID, embedMsg) //if member writes unknown message give them fhelp embed
		}
	}
}
