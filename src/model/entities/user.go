package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	BaseEntity
	ID   primitive.ObjectID `bson:"_id" json:"id"`
	Name string             `bson:"name" json:"name"`
	Age  int                `bson:"age" json:"age"`
}
