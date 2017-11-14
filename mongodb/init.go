package mongodb

import "os"

var connectionString string

func init() {
	connectionString = os.Getenv("MONGODB_CONNECTION_STRING")
	if len(connectionString) == 0 {
		// for local
		connectionString = "mongodb://admin:admin@ds249325.mlab.com:49325/werpas"
	}
}

