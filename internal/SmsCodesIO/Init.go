package SmsCodesIO

func Init() *Session {
	session := &Session{
		ApiKey:      "dd6cdc3f-bdd2-4250-8f8a-220390c72f6c",
		Country:     "UK",
		ServiceID:   "462f7a96-98e9-44a5-9407-47d3104519bd",
		SerciceName: "Foodora",
		Number:      "",
		SecurityID:  "",
	}
	getNumber(session)
	return session
}
