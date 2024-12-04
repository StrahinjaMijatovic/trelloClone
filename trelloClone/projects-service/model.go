package main

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	Name       string             `json:"name" bson:"name"`
	EndDate    time.Time          `json:"endDate" bson:"end_date"`
	MinMembers int                `json:"minMembers" bson:"min_members"`
	MaxMembers int                `json:"maxMembers" bson:"max_members"`
	ManagerID  string             `json:"managerId" bson:"manager_id"`
	Members    []string           `json:"members" bson:"members"`
	CreatedAt  time.Time          `json:"createdAt" bson:"created_at"`
	UpdatedAt  time.Time          `json:"updatedAt" bson:"updated_at"`
}
