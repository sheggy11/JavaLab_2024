package core

import "go.mongodb.org/mongo-driver/bson/primitive"

type CV struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name,omitempty"`
	ActivityType string             `bson:"activityType,omitempty"`
	Summary      string             `bson:"summary,omitempty"`
}
