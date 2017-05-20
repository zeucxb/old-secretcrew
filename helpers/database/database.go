package database

import (
	"os"

	log "github.com/Sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"
)

var databaseName string
var Session *mgo.Session

func StartDB() *mgo.Session {
	strcon := getStrCon()

	var err error

	Session, err = mgo.Dial(strcon)
	if err != nil {
		panic(err)
	}

	return Session
}

func CloseDB() {
	Session.Close()
}

func getStrCon() (strcon string) {
	strcon = "mongodb://localhost"

	if mongoURI := os.Getenv("MONGODB_URI"); mongoURI != "" {
		strcon = mongoURI
	}

	return
}

// DB return a mongodb database connection
func DB(dbName string) {
	Session.DB(dbName)
}

// SetDB set a mongodb database
func SetDB(dbName string) {
	databaseName = dbName
}

// Collection return a mongodb collection
func Collection(collectionName string, index mgo.Index) (c *mgo.Collection) {
	c = Session.DB(databaseName).C(collectionName)

	if err := c.EnsureIndex(index); err != nil {
		log.Warn(err)
	}

	return
}
