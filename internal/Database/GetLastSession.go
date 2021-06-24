package Database

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"smsbot/internal/SmsCodesIO"
)
func GetLastSession(discordID string) *SmsCodesIO.LastSession {
	dbsession := getDatabase(false, &DatabaseSession{})

	data := &SmsCodesIO.UserData{}

	err := dbsession.collectionPtr.FindOne(nil, bson.M{"discord_id":discordID}).Decode(data)
	if err != nil{
		fmt.Println(err)
	}

	lastSession := &SmsCodesIO.LastSession{
		Apikey:     data.LastSession.Apikey,
		Country:     data.LastSession.Country,
		ServiceID:   data.LastSession.ServiceID,
		ServiceName: data.LastSession.ServiceName,
		Number:      data.LastSession.Number,
		SecurityID:  data.LastSession.SecurityID,
	}
	return lastSession
}
