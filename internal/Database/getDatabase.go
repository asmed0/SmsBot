package Database

var passData *DatabaseSession

func getDatabase(init bool, data *DatabaseSession) *DatabaseSession {
	switch init {
	case true:
		passData = data
		return passData
	case false:
		return passData
	}
	return passData //cannot happen
}
