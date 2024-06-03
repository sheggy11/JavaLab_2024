package core

import "go.mongodb.org/mongo-driver/bson/primitive"

type Job struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Name        string             `bson:"name,omitempty" json:"name"`
    Experience  int32              `bson:"experience,omitempty" json:"experience"`
    Skills      []string           `bson:"skills,omitempty" json:"skills"`
    Description string             `bson:"description,omitempty" json:"description"`
    Company     string             `bson:"company,omitempty" json:"company"`
}
