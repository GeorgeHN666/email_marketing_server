package db

import (
	"context"
	"time"

	"github.com/GeorgeHN/email-backend/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *DB) InsertSchedules(c models.Schedule) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*45)
	defer cancel()

	db := s.Conn.Database(s.config.Database).Collection("schedules")

	c.ID = primitive.NewObjectID()
	c.Issued = time.Now().Local().Unix()
	c.Status = Active

	_, err := db.InsertOne(ctx, c)
	if err != nil {
		return err
	}

	return nil
}

// Delete Schedules
func (s *DB) DeleteSchedules(i string) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*45)
	defer cancel()

	db := s.Conn.Database(s.config.Database).Collection("schedules")

	id, _ := primitive.ObjectIDFromHex(i)

	filter := bson.M{
		"_id": bson.M{"$eq": id},
	}

	_, err := db.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (s *DB) GetSchedules(i string) ([]*models.Schedule, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*45)
	defer cancel()

	db := s.Conn.Database(s.config.Database).Collection("schedules")

	id, _ := primitive.ObjectIDFromHex(i)

	var result []*models.Schedule

	filter := bson.M{
		"campaing_id": bson.M{"$eq": id},
	}

	op := options.Find().SetSort(bson.M{"issued": -1})

	cursor, err := db.Find(ctx, filter, op)
	if err != nil {
		return result, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {

		var d models.Schedule

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

func (s *DB) GetSchedule(i string) (*models.Schedule, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*45)
	defer cancel()

	db := s.Conn.Database(s.config.Database).Collection("schedules")

	id, _ := primitive.ObjectIDFromHex(i)

	filter := bson.M{
		"_id": bson.M{"$eq": id},
	}

	var res models.Schedule

	err := db.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil

}
