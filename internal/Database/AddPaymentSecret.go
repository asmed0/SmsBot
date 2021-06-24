package Database

import (
	"go.mongodb.org/mongo-driver/bson"
)

func AddPaymentSecret(discordID string, secret string){
	dbsession := getDatabase(false, &DatabaseSession{})
	dbsession.collectionPtr.FindOneAndUpdate(nil, bson.M{"discord_id": discordID},
		bson.D{
			{"$set", bson.D{
				{"payment_secret", secret}}}})
}
