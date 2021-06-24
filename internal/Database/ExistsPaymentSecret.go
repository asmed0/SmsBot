package Database

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func ExistsPaymentSecret(secret string) bool{
	dbsession := getDatabase(false, &DatabaseSession{})
	exists := dbsession.collectionPtr.FindOne(nil, bson.M{"payment_secret": secret})
	if exists.Err() == mongo.ErrNoDocuments{
		return false
	}
	return true
}
