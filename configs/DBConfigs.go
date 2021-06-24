package configs

import "os"

func DBConfigs() *DBStruct {
	config := &DBStruct{
		Uri:        os.Getenv("db_uri"),
		User:       os.Getenv("db_user"),
		Pass:       os.Getenv("db_pass"),
		Database:   os.Getenv("db_database"),
		Collection: os.Getenv("db_collection"),
	}
	return config
}
