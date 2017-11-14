package mongodb

import "os"

var connectionString string
var database string
var collection string

func init() {
	connectionString = os.Getenv("MONGODB_CONNECTION_STRING")
	if len(connectionString) == 0 {
		// for local
		connectionString = "mongodb://admin:admin@ds249325.mlab.com:49325/werpas"
	}

	database = os.Getenv("MONGODB_DB")
	if len(database) == 0 {
		// for local
		database = "werpas"
	}

	collection = os.Getenv("MONGODB_COLLECTION")
	if len(collection) == 0 {
		// for local
		collection = "event"
	}
}

