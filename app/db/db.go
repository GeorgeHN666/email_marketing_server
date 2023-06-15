package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Authorized    = 0 // Used for codes, auths, etc
	Active        = 1 // Used for campaings, client status.
	Done          = 2 // Used for schedules
	Pending       = 3 // used for schedules
	NotAuthorized = 4 // Used for codes, auths, etc
	NotActive     = 5 // Used for campaings, client status.
	Error         = 6 // used for schedules and any other error
)

type DB struct {
	config Config
	Conn   *mongo.Client
}

func NewDBConn(config Config) *DB {

	c, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.URI))
	if err != nil {
		return nil
	}

	err = c.Ping(context.TODO(), nil)
	if err != nil {
		// SEND EMAIL INFORMING THE ERROR IN THE DATABASE
		return nil
	}

	return &DB{
		config: config,
		Conn:   c,
	}
}

// Temporary
