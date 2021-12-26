package handler

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) initArticlesRoutes(r *gin.Engine) {
	articles := r.Group("/articles")
	{
		articles.GET("", h.getAllArticles)
		articles.GET("/:id", h.getArticleByID)
		articles.POST("/:id", h.createArticle)
	}
}

func (h *Handler) getAllArticles(c *gin.Context) {
	articles := h.services.Articles.GetAllArticles(c)

	c.JSON(200, articles)
}

func (h *Handler) getArticleByID(c *gin.Context) {
	id, err := parseIdFromPath(c, "_id")
	if err != nil {
		log.Println(err)
	}

	article, err := h.services.Articles.GetArticleByID(c, id)
	if err != nil {
		log.Println()
	}

	c.JSON(200, article)
}

func (h *Handler) createArticle(c *gin.Context) {

}

func parseIdFromPath(c *gin.Context, param string) (primitive.ObjectID, error) {
	idParam := c.Param(param)
	if idParam == "" {
		return primitive.ObjectID{}, errors.New("empty id param")
	}

	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return primitive.ObjectID{}, errors.New("invalid id param")
	}

	return id, nil
}
