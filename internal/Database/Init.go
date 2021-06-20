package Database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var err error

func Init() {
	dbsession := &DatabaseSession{
		uri:        "cluster0.zc0f8.mongodb.net",
		user:       "dbUser",
		pass:       "arfiestysexu",
		database:   "Database",
		collection: "SmsBot",
	}

	dbsession.client, err = mongo.NewClient(options.Client().ApplyURI("mongodb+srv://" + dbsession.user +
		":" + dbsession.pass +
		"@" + dbsession.uri +
		"/" + dbsession.database + "?retryWrites=true&w=majority"))

	if err != nil {
		log.Fatal(err)
	}

	dbsession.ctx, _ = context.WithTimeout(context.Background(), 1000*time.Second)
	err = dbsession.client.Connect(dbsession.ctx)
	if err != nil {
		log.Fatal(err)
	}

	//defer dbsession.client.Disconnect(dbsession.ctx) //kills the server, will reroute later
	err = dbsession.client.Ping(dbsession.ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	dbsession.collectionPtr = dbsession.client.Database(dbsession.database).Collection(dbsession.collection)
	//fmt.Println(dbsession.client.Database(dbsession.database).Collection(dbsession.collection).FindOne(dbsession.ctx, bson.M{}).DecodeBytes())
	getDatabase(true, dbsession) //initializing our getdb variables
}
