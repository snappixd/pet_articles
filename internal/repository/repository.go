package repository

import (
	"articles/internal/models"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Articles interface {
	Create(ctx context.Context, article models.Article) (primitive.ObjectID, error)
	Delete(ctx context.Context, articleId primitive.ObjectID) error
	GetAllArticles(ctx context.Context) []models.Article
	GetArticleByID(ctx context.Context, articleId primitive.ObjectID) (models.Article, error)
}

type Repositories struct {
	Articles Articles
}

func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		Articles: NewArticlesRepo(db),
	}
}
