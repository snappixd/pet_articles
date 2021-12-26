package service

import (
	"articles/internal/models"
	"articles/internal/repository"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Articles interface {
	Create(ctx context.Context, input models.Article) (primitive.ObjectID, error)
	Delete(ctx context.Context, articleId primitive.ObjectID) error
	GetAllArticles(ctx context.Context) []models.Article
	GetArticleByID(ctx context.Context, articleId primitive.ObjectID) (models.Article, error)
}

type Services struct {
	Articles Articles
}

type Deps struct {
	Repos *repository.Repositories
}

func NewServices(deps Deps) *Services {
	articlesServices := NewArticlesService(deps.Repos.Articles)
	return &Services{
		Articles: articlesServices,
	}
}
