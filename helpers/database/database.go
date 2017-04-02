package database

import (
	"os"

	mgo "gopkg.in/mgo.v2"
)

var databaseName string

func getSession() *mgo.Session {
	strcon := "mongodb://localhost"

	if mongoURI := os.Getenv("MONGODB_URI"); mongoURI != "" {
		strcon = mongoURI
	}

	// Connect to our mongodb
	s, err := mgo.Dial(strcon)
	if err != nil {
		panic(err)
	}

	return s
}

// DB return a mongodb database connection
func DB(dbName string) {
	getSession().DB(databaseName)
}

// SetDB set a mongodb database
func SetDB(dbName string) {
	databaseName = dbName
}

// Collection return a mongodb collection
func Collection(collectionName string) *mgo.Collection {
	return getSession().DB(databaseName).C(collectionName)
}
