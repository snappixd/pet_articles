package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Article struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Anons       string             `json:"anons" bson:"anons"`
	Description string             `json:"description" bson:"description"`
	Author      string             `json:"author" bson:"author"`
}
