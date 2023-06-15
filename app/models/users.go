package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Admin struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id"`
	Name         string             `json:"name" bson:"name"`
	Email        string             `json:"email" bson:"email"`
	Password     string             `json:"password" bson:"password"`
	Code         int                `json:"code" bson:"code"`
	SecurityCode int                `json:"security_code" bson:"security_code"`
	ValidCode    int                `json:"valid_code" bson:"valid_code"`
}

type Client struct {
	ID      primitive.ObjectID `json:"_id" bson:"_id"`
	Name    string             `json:"name" bson:"name"`
	Email   string             `json:"contact" bson:"contact"`
	Company string             `json:"company"  bson:"company"`
	Since   time.Time          `json:"since"  bson:"since"`
	Status  int                `json:"status"  bson:"status"`
}

type Audience struct {
	First string `json:"first"`
	Last  string `json:"Last"`
	Email string `json:"email"`
}

type Campaing struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id"`
	ClientID     primitive.ObjectID `json:"client_id" bson:"client_id"`
	ClientName   string             `json:"client_name" bson:"client_name"`
	CompanyName  string             `json:"company_name" bson:"company_name"`
	Audience     string             `json:"audience" bson:"audience"`
	CampaingName string             `json:"campaing_name" bson:"campaing_name"`
	StartDate    time.Time          `json:"start_date" bson:"start_date"`
	EndDate      time.Time          `json:"end_date" bson:"end_date"`
	Issued       time.Time          `json:"issued" bson:"issued"`
	Status       int                `json:"status" bson:"status"`
}

type Schedule struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id"`
	Name         string             `json:"name"  bson:"name"`
	CampaingID   primitive.ObjectID `json:"campaing_id" bson:"campaing_id"`
	CampaingName string             `json:"campaing_name"  bson:"campaing_name"`
	CompanyName  string             `json:"company_name"  bson:"company_name"`
	TimeSet      time.Time          `json:"time_set" bson:"time_set"`
	Template     string             `json:"template" bson:"template"`
	Subject      string             `json:"subject" bson:"subject"`
	PrevText     string             `json:"prev_text" bson:"prev_text"`
	Issued       time.Time          `json:"issued" bson:"issued"`
	Status       int                `json:"status" bson:"status"`
}

type Visit struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	CampaingID primitive.ObjectID `json:"campaing_id" bson:"campaing_id"`
	ScheduleID primitive.ObjectID `json:"schedule_id" bson:"schedule_id"`
	Count      int                `json:"count"  bson:"count"`
	Regs       []Reg              `json:"regs"  bson:"regs"`
}
type Reg struct {
	TimeVisited time.Time `json:"time_visited"  bson:"time_visited"`
}
