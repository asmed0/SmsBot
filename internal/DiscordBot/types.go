package DiscordBot

type DiscordData struct {
	Token    string   `json:"token"`
	Prefix   string   `json:"prefix"`
	Commands [][]string `json:"commands"`
}
