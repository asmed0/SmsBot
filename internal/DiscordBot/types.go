package DiscordBot

type DiscordData struct {
	Token    string   `json:"token"`
	AppID string `json:"appid"`
	Prefix   string   `json:"prefix"`
	Commands [][]string `json:"commands"`
}
