package Database

import (
	"go.mongodb.org/mongo-driver/bson"
	"smsbot/internal/FiveSim"
	"smsbot/internal/SmsCodesIO"
)

var userData SmsCodesIO.UserData

func UpdateSession(discordID string, lastSession *FiveSim.FiveSimSession, dispose bool) string {
	dbsession := getDatabase(false, &DatabaseSession{})
	dbsession.collectionPtr.FindOne(nil, bson.M{"discord_id": discordID}).Decode(&userData)
	if userData.ID == ""{
		dbsession.collectionPtr.InsertOne(nil, bson.M{"discord_id": discordID}) //first inserting
		dbsession.collectionPtr.FindOneAndUpdate(nil, bson.M{"discord_id": discordID},
			bson.D{
				{"$set", bson.D{
					{"balance", 0}}}})
		return "zerobal" //no balance but initialized user
	}

	if userData.Balance <= 0{
		return "zerobal"
	}

	dbsession.collectionPtr.FindOneAndUpdate(dbsession.ctx, bson.M{"discord_id": discordID},
		bson.D{
			{"$set", bson.D{
				{"last_session", bson.D{
					{"apikey", lastSession.ApiKey},
					{"id", lastSession.ID},
					{"phone", lastSession.Phone},
					{"operator", lastSession.Operator},
					{"product", lastSession.Product},
					{"price", lastSession.Price},
					{"status", lastSession.Status},
					{"expires", lastSession.Expires},
					{"sms", lastSession.Sms},
					{"created_at", lastSession.CreatedAt},
					{"country", lastSession.Country},
					{"is_disposed", dispose}}}}}}).DecodeBytes()
	return lastSession.Phone
}
