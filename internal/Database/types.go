package Database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type DatabaseSession struct {
	uri           string
	user          string
	pass          string
	database      string
	collection    string
	collectionPtr *mongo.Collection
	ctx           context.Context
	client        *mongo.Client
}

