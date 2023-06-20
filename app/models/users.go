package models

import (
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
	Since   int64              `json:"since"  bson:"since"`
	Status  int                `json:"status"  bson:"status"`
	Plan    string             `json:"plan"  bson:"plan"`
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
	StartDate    int64              `json:"start_date" bson:"start_date"`
	EndDate      int64              `json:"end_date" bson:"end_date"`
	Issued       int64              `json:"issued" bson:"issued"`
	Status       int                `json:"status" bson:"status"`
}

type Schedule struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id"`
	Name         string             `json:"name"  bson:"name"`
	CampaingID   primitive.ObjectID `json:"campaing_id" bson:"campaing_id"`
	CampaingName string             `json:"campaing_name"  bson:"campaing_name"`
	CompanyName  string             `json:"company_name"  bson:"company_name"`
	TimeSet      int64              `json:"time_set" bson:"time_set"`
	Audience     string             `json:"audience"  bson:"audience"`
	Template     string             `json:"template" bson:"template"`
	Subject      string             `json:"subject" bson:"subject"`
	PrevText     string             `json:"prev_text" bson:"prev_text"`
	Issued       int64              `json:"issued" bson:"issued"`
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
	TimeVisited int64 `json:"time_visited"  bson:"time_visited"`
}

// TODOOO

// We need to show the schedules base on the day selected
//  Upload Template
//  Upload Audience
// Make the timers and launch go functions to be triggered
// Configure Middlewares for statistics
