package Database

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
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

	err = dbsession.client.Connect(nil)
	if err != nil {
		log.Fatal(err)
	}

	//defer dbsession.client.Disconnect(dbsession.ctx) //kills the server, will reroute later
	err = dbsession.client.Ping(nil, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	dbsession.collectionPtr = dbsession.client.Database(dbsession.database).Collection(dbsession.collection)
	//fmt.Println(dbsession.client.Database(dbsession.database).Collection(dbsession.collection).FindOne(dbsession.ctx, bson.M{}).DecodeBytes())
	getDatabase(true, dbsession) //initializing our getdb variables
}
