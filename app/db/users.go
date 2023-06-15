package db

import (
	"context"
	"time"

	"github.com/GeorgeHN/email-backend/app/models"
	"github.com/GeorgeHN/email-backend/app/serializers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (s *DB) GetClients() ([]*models.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*45)
	defer cancel()

	db := s.Conn.Database(s.config.Database).Collection("clients")

	var result []*models.Client

	filter := bson.M{}

	op := options.Find().SetSort(bson.M{"since": -1})

	cursor, err := db.Find(ctx, filter, op)
	if err != nil {
		return result, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {

		var d models.Client

		err := cursor.Decode(&d)
		if err != nil {
			return result, err
		}

		result = append(result, &d)

	}

	err = cursor.Err()
	if err != nil {
		return nil, err
	}

	return result, nil

}

func (s *DB) GetClient(id string) (*models.Client, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*45)
	defer cancel()

	db := s.Conn.Database(s.config.Database).Collection("clients")

	i, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{
		"_id": bson.M{"$eq": i},
	}

	var res models.Client

	err := db.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
