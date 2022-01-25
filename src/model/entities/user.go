package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	BaseEntity
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
	Age  int                `bson:"age"`
}
