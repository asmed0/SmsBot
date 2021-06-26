package Database

import (
	"go.mongodb.org/mongo-driver/bson"
	"smsbot/internal/SmsCodesIO"
)

var userData SmsCodesIO.UserData

func UpdateSession(discordID string, lastSession *SmsCodesIO.Session, dispose bool) string {
	dbsession := getDatabase(false, &DatabaseSession{})
	dbsession.collectionPtr.FindOne(nil, bson.M{"discord_id": discordID}).Decode(&userData)
	if userData.ID == ""{
		dbsession.collectionPtr.InsertOne(nil, bson.M{"discord_id": discordID}) //first inserting
		dbsession.collectionPtr.FindOneAndUpdate(nil, bson.M{"discord_id": discordID},
			bson.D{
				{"$set", bson.D{
					{"balance", 0}}}})
		return "zerobal" //no balance initialized user
	}

	if userData.Balance == 0{
		return "zerobal"
	}

	dbsession.collectionPtr.FindOneAndUpdate(dbsession.ctx, bson.M{"discord_id": discordID},
		bson.D{
			{"$set", bson.D{
				{"last_session", bson.D{
					{"apikey", lastSession.ApiKey},
					{"country", lastSession.Country},
					{"service_id", lastSession.ServiceID},
					{"service_name", lastSession.SerciceName},
					{"number", lastSession.Number},
					{"security_id", lastSession.SecurityID},
					{"is_disposed", dispose}}}}}}).DecodeBytes()
	return lastSession.Number
}
