package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) initArticlesRoutes(r *gin.Engine) {
	articles := r.Group("/")
	{
		articles.GET("/", h.getAllArticles)
		articles.GET("/getArticle/:id", h.getArticleByID)
		articles.GET("/deleteArticle/:id", h.deleteArticle)
		articles.POST("/createArticle/:id", h.createArticle)
	}
}

func (h *Handler) getAllArticles(c *gin.Context) {
	// if c.URL.Path != "/" {
	// 	log.Fatal("err")
	// }

	articles := h.services.Articles.GetAllArticles(c)

	c.HTML(http.StatusOK, "home.page.html", gin.H{
		"articles": articles,
	})
}

func (h *Handler) getArticleByID(c *gin.Context) {
	id, err := parseIdFromPath(c, "id")
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
	// id := primitive.NewObjectID()
}

func (h *Handler) deleteArticle(c *gin.Context) {
	id, err := parseIdFromPath(c, "id")
	if err != nil {
		log.Println(err)
	}

	err = h.services.Articles.Delete(c, id)

	c.Redirect(http.StatusMovedPermanently, "/articles")
	c.Abort()
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
