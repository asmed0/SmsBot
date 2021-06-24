package Database

import (
	"go.mongodb.org/mongo-driver/bson"
	"smsbot/internal/SmsCodesIO"
)

func GetBalance(discordID string)int{
	dbsession := getDatabase(false, &DatabaseSession{})
	data := &SmsCodesIO.UserData{}
	dbsession.collectionPtr.FindOne(nil, bson.M{"discord_id": discordID} ).Decode(data)
	return data.Balance
}

