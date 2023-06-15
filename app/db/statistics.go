package db

import (
	"context"
	"time"

	"github.com/GeorgeHN/email-backend/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *DB) InsertEmailVisit(i string, reg *models.Reg) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*45)
	defer cancel()

	db := s.Conn.Database(s.config.Database).Collection("statistics")

	id, _ := primitive.ObjectIDFromHex(i)

	filter := bson.M{
		"_id": bson.M{"$eq": id},
	}

	change := bson.M{
		"$inc":  bson.M{"count": 1},
		"$push": bson.M{"regs": reg},
	}

	_, err := db.UpdateOne(ctx, filter, change)
	if err != nil {
		return err
	}
	return nil
}
