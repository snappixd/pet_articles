package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Article struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Anons       string             `json:"anons" bson:"anons"`
	Description string             `json:"description" bson:"description"`
	Author      *Author            `json:"author" bson:"author"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
}

// type CreateArticleInput struct {
// 	Title       string  `json:"title" bson:"title"`
// 	Anons       string  `json:"anons" bson:"anons"`
// 	Description string  `json:"description" bson:"description"`
// 	Author      *Author `json:"author" bson:"author"`
// }
