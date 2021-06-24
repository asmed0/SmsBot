package Database

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
	exists := session.collectionPtr.FindOne(nil, bson.M{"discord_id": discordID})
	if exists.Err() == mongo.ErrNoDocuments{
		session.collectionPtr.FindOneAndUpdate(nil, bson.M{"discord_id": discordID},
			bson.D{
				{"$set", bson.D{
					{"balance", 0}}}})
	}
	session.collectionPtr.FindOne(nil, bson.M{"discord_id": discordID} ).Decode(data)
	return data.Balance
}
