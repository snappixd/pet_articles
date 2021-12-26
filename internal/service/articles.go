package service

import (
	"articles/internal/models"
	"articles/internal/repository"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ArticlesService struct {
	repo repository.Articles
}

func NewArticlesService(repo repository.Articles) *ArticlesService {
	return &ArticlesService{
		repo: repo,
	}
}

func (s *ArticlesService) Create(ctx context.Context, article models.Article) (primitive.ObjectID, error) {
	return s.repo.Create(ctx, article)
}

func (s *ArticlesService) Delete(ctx context.Context, articleId primitive.ObjectID) error {
	return s.repo.Delete(ctx, articleId)
}

func (s *ArticlesService) GetAllArticles(ctx context.Context) []models.Article {
	return s.repo.GetAllArticles(ctx)
}

func (s *ArticlesService) GetArticleByID(ctx context.Context, articleId primitive.ObjectID) (models.Article, error) {
	return s.repo.GetArticleByID(ctx, articleId)
}
