package Database

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddPaymentSecret(discordID string, secret string){
	dbsession := getDatabase(false, &DatabaseSession{})
	exists := dbsession.collectionPtr.FindOne(nil, bson.M{"discord_id": discordID})
	if exists.Err() == mongo.ErrNoDocuments{
		dbsession.collectionPtr.InsertOne(nil, bson.M{"discord_id": discordID})
	}

	dbsession.collectionPtr.FindOneAndUpdate(nil, bson.M{"discord_id": discordID},
		bson.D{
			{"$set", bson.D{
				{"payment_secret", secret}}}})
}
