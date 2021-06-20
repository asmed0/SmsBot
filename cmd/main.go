package cmd

import (
	"SmsBot/internal/DiscordBot"
)

func main() {
	discordData := &DiscordData{
		Token:    "ODM0ODIwMDg3MDg1MDA2ODkw.YIGcyg.dmEPblcFyqOKT06Sog4Ary8zDGM",
		Prefix:   "!",
		Commands: []string{"food", "balance", "topup", "help"},
	}

	DiscordBot.Init(discordData)

}
