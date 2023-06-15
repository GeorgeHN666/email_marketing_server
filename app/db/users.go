package db

import (
	"context"
	"time"

	"github.com/GeorgeHN/email-backend/app/models"
	"github.com/GeorgeHN/email-backend/app/serializers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *DB) InsertAdmin(ad models.Admin) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*45)
	defer cancel()

	db := s.Conn.Database(s.config.Database).Collection("admin")

	ad.ID = primitive.NewObjectID()
	ad.SecurityCode, _ = serializers.GenerateNumCode()
	ad.ValidCode = NotAuthorized
	ad.Password, _ = serializers.Hash(ad.Password)

	_, err := db.InsertOne(ctx, ad)
	if err != nil {
		return err
	}

	return nil
}

func (s *DB) InsertClient(c models.Client) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*45)
	defer cancel()

	db := s.Conn.Database(s.config.Database).Collection("clients")

	c.ID = primitive.NewObjectID()
	c.Since = time.Now()
	c.Status = Active

	_, err := db.InsertOne(ctx, c)
	if err != nil {
		return err
	}

	return nil
}

func (s *DB) DeleteClient(id primitive.ObjectID) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*45)
	defer cancel()

	db := s.Conn.Database(s.config.Database).Collection("clients")

	filter := bson.M{
		"_id": bson.M{"$eq": id},
	}

	_, err := db.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
