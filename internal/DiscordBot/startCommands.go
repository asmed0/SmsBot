package DiscordBot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"smsbot/internal/Database"
	"smsbot/internal/FiveSim"
	"smsbot/internal/Topup"
	"smsbot/internal/tools"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

var s *discordgo.Session

// Important note: call every command in order it's placed in the example.

var (

	embedMsg = &discordgo.MessageEmbed{
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
	boltMsg = &discordgo.MessageEmbed{
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
	infoEmbed = &discordgo.MessageEmbed{
		Title:  "+1 == Kanada",
		Description: "**Landskoden +1 är för Kanada**",
		Fields: []*discordgo.MessageEmbedField{},
		Provider: &discordgo.MessageEmbedProvider{
			URL:  "https://cdn.discordapp.com/icons/806511362251030558/244ed44d2ab37a59e37bb775de0d8fcb.png?size=256",
			Name: "SlotTalk SMSBOT",
		},
		Color: 16776960, //yellow color
	}

	firstBtn = discordgo.ActionsRow{
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
	boltBtn = discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{
			discordgo.Button{
				Emoji: discordgo.ComponentEmoji{
					Name: "🎢",
				},
				Label:    "Request Bolt number!",
				Style:    discordgo.SuccessButton,
				CustomID: "bolt",
			},
		},
	}
	componentsHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"imready": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			getNumber(embedMsg, i.Interaction.User.ID, "other", -2)
			if embedMsg.Color != 15158332 {
				err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: embedMsg.Title,
						Flags:   1 << 6,
						Embeds:  []*discordgo.MessageEmbed{embedMsg,infoEmbed},
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
					fmt.Println(err)
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
										Label: "Click here to topup!",
										Style: discordgo.LinkButton,
										URL:   "https://checkout.stripe.com/pay/" + Topup.CreateCheckoutSession(i.Interaction.User.ID, 10),
									},
								},
							},
						},
					},
				})
				if err != nil {
					fmt.Println(err)
				}
			}
		},
		"getcode": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			returnedCode := FiveSim.GetSms(Database.GetLastSession(i.Interaction.User.ID))
			if returnedCode == "Err" {
				embedMsg.Title = "Message not received yet, try again in a moment"
				embedMsg.Description = ""
				embedMsg.Color = 15158332 //red color
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
					fmt.Println(err)
				}
			} else if returnedCode == "ProviderErr" {
				embedMsg.Title = "Seems like you took to long to use the number, or our provider had an error."
				//embedMsg.Description = "If this is an issue on our end, contact Sithed"
				embedMsg.Color = 15158332 //red color
				//go Database.UpdateBalance(1, i.Interaction.User.ID)
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
					fmt.Println(err)
				}
			} else {
				embedMsg.Title = returnedCode
				embedMsg.Description = "Thank you for using our services, above is an auto-generated 200/70 voucher"
				embedMsg.Color = 3066993 //green color

				lastSession := Database.GetLastSession(i.Interaction.User.ID)
				if returnedCode != "Err" && returnedCode != "ProverErr" {
					lastSession.Sms = append(lastSession.Sms, FiveSim.SmsSlice{Text: returnedCode})
				}
				Database.UpdateSession(i.Interaction.User.ID, &FiveSim.FiveSimSession{
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

				err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: tools.DiscountGen(),
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
					fmt.Println(err)
				}
			}
		},
		"getbalance": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			embedMsg.Title = strconv.Itoa(Database.GetBalance(i.Interaction.User.ID)) + " Tokens left"
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
				fmt.Println(err)
			}
		},

		"bolt": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			getNumber(embedMsg, i.Interaction.User.ID, "bolt", -2)
			if embedMsg.Color != 15158332 {
				err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: embedMsg.Title,
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
					fmt.Println(err)
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
										Label: "Click here to topup!",
										Style: discordgo.LinkButton,
										URL:   "https://checkout.stripe.com/pay/" + Topup.CreateCheckoutSession(i.Interaction.User.ID, 10),
									},
								},
							},
						},
					},
				})
				if err != nil {
					fmt.Println(err)
				}
			}
		},
	}
)

func startCommands(data *DiscordData) {
	var err error
	s, err = discordgo.New("Bot " + data.Token)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}

	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("Command session is online!")
	})
	// Components are part of interactions, so we register InteractionCreate handler
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := componentsHandlers[i.MessageComponentData().CustomID]; ok {
			embedMsg = embedCleaner() // cleaning the embed since previous usage
			h(s, i)
			go sendLogs(i.Interaction.User.Username, embedMsg, s, i.Interaction.User.ID)
		}
	})

	err = s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}
	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Graceful shutdown")
}
