package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	BaseEntity `bson:",omitempty,inline"`
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string             `bson:"name"`
	Age        int                `bson:"age"`
}
