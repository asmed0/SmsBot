package DiscordBot

var passData *DiscordData

func getConfig(init bool, data *DiscordData) *DiscordData {
	switch init {
	case true:
		passData = data
		return passData
	case false:
		return passData
	}
	return passData //cannot happen
}
