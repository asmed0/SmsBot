package DiscordBot

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func newMember(s *discordgo.Session, m *discordgo.GuildMemberUpdate){
	directMessage, err := s.UserChannelCreate(m.User.ID)
	if err != nil {
		log.Panic(err)
	}
	s.ChannelMessageSend(directMessage.ID,"hello")

}
