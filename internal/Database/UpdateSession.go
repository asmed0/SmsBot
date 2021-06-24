package Database

import (
	"go.mongodb.org/mongo-driver/bson"
	"smsbot/internal/SmsCodesIO"
)

func UpdateSession(discordID string, lastSession *SmsCodesIO.Session) string {
	dbsession := getDatabase(false, &DatabaseSession{})
	dbsession.collectionPtr.FindOneAndUpdate(dbsession.ctx, bson.M{"discord_id": discordID},
		bson.D{
			{"$set", bson.D{
				{"last_session", bson.D{
					{"apikey", lastSession.ApiKey},
					{"country", lastSession.Country},
					{"service_id", lastSession.ServiceID},
					{"service_name", lastSession.SerciceName},
					{"number", lastSession.Number},
					{"security_id", lastSession.SecurityID}}}}}}).DecodeBytes()
	return lastSession.Number
}
