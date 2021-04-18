package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Contact struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Address     string             `bson:"address"`
	PhoneNumber []int              `bson:"phone_number"`
	Owners      []string           `bson:"owners"`
}
