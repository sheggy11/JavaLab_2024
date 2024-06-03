package core

import "go.mongodb.org/mongo-driver/bson/primitive"

type Job struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Skills      []string           `bson:"skills,omitempty"`
	Revenue     uint               `bson:"revenue,omitempty"`
	CompanyName string             `bson:"companyName,omitempty"`
}
