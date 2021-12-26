package repository

import (
	"articles/internal/models"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ArticlesRepo struct {
	db *mongo.Collection
}

func NewArticlesRepo(db *mongo.Database) *ArticlesRepo {
	return &ArticlesRepo{
		db: db.Collection(articlesCollection),
	}
}

func (r *ArticlesRepo) Create(ctx context.Context, article models.Article) (primitive.ObjectID, error) {
	article.ID = primitive.NewObjectID()

	_, err := r.db.InsertOne(ctx, article)

	return article.ID, err
}

func (r *ArticlesRepo) Delete(ctx context.Context, articleId primitive.ObjectID) error {
	_, err := r.db.DeleteOne(ctx, bson.M{"_id": articleId})

	return err
}

func (r *ArticlesRepo) GetAllArticles(ctx context.Context) []models.Article {
	var articles []models.Article

	cursor, err := r.db.Find(ctx, bson.D{})
	if err != nil {
		log.Println(err)
	}

	for cursor.Next(context.TODO()) {
		var article models.Article

		err := cursor.Decode(&article)
		if err != nil {
			log.Println(err)
		}

		articles = append(articles, article)
	}

	if err := cursor.Err(); err != nil {
		log.Println(err)
	}

	cursor.Close(context.TODO())

	return articles
}

func (r *ArticlesRepo) GetArticleByID(ctx context.Context, articleId primitive.ObjectID) (models.Article, error) {
	var article models.Article

	err := r.db.FindOne(ctx, bson.M{"_id": articleId}).Decode(&article)

	return article, err
}
