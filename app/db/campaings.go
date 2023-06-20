package db

import (
	"context"
	"time"

	"github.com/GeorgeHN/email-backend/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *DB) InsertCampaing(c models.Campaing) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*45)
	defer cancel()

	db := s.Conn.Database(s.config.Database).Collection("campaings")

	c.Issued = time.Now().Local().Unix()
	c.ID = primitive.NewObjectID()
	c.Status = Active

	_, err := db.InsertOne(ctx, c)
	if err != nil {
		return err
	}

	return nil
}

// Delete campaings
func (s *DB) DeleteCampaing(id string) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*45)
	defer cancel()

	db := s.Conn.Database(s.config.Database).Collection("campaings")

	i, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{
		"_id": bson.M{"$eq": i},
	}

	_, err := db.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (s *DB) GetCampaings(i string) ([]*models.Campaing, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*45)
	defer cancel()

	db := s.Conn.Database(s.config.Database).Collection("campaings")

	id, _ := primitive.ObjectIDFromHex(i)

	var result []*models.Campaing

	filter := bson.M{
		"client_id": bson.M{"$eq": id},
	}

	op := options.Find().SetSort(bson.M{"issued": -1})

	cursor, err := db.Find(ctx, filter, op)
	if err != nil {
		return result, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {

		var d models.Campaing

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

func (s *DB) GetCampaing(i string) (*models.Campaing, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*45)
	defer cancel()

	db := s.Conn.Database(s.config.Database).Collection("campaings")

	id, _ := primitive.ObjectIDFromHex(i)

	filter := bson.M{
		"_id": bson.M{"$eq": id},
	}

	var res models.Campaing

	err := db.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
