package Database

import (
	"go.mongodb.org/mongo-driver/bson"
	"smsbot/internal/SmsCodesIO"
)

func UpdateBalance(num int, discordID string){
	dbsession := getDatabase(false, &DatabaseSession{})
	prevBalance := getBalanceSafe(dbsession, discordID)
	dbsession.collectionPtr.FindOneAndUpdate(nil, bson.M{"discord_id": discordID},
		bson.D{
			{"$set", bson.D{
				{"balance", prevBalance + num}}}})
}

func getBalanceSafe(session *DatabaseSession, discordID string)int{
	data := &SmsCodesIO.UserData{}
	session.collectionPtr.FindOne(nil, bson.M{"discord_id": discordID} ).Decode(data)
	return data.Balance
}
