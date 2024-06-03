package core

import "go.mongodb.org/mongo-driver/bson/primitive"

type CV struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Name   string             `bson:"candidateName,omitempty"`
	Age    uint               `bson:"candidateAge,omitempty"`
	Skills []string           `bson:"skills,omitempty"`
}
