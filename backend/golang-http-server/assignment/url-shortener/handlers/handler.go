package handlers

import (
	"log"
	"net/http"

	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/entity"
	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/repository"

	"github.com/gin-gonic/gin"
)

type URLHandler struct {
	repo *repository.URLRepository
}

func NewURLHandler(repo *repository.URLRepository) URLHandler {
	return URLHandler{
		repo: repo,
	}
}

func (h *URLHandler) Get(c *gin.Context) {
	// TODO: answer here

	// ambil param
	path := c.Param("path")

	// cek ke get repo
	url, err := h.repo.Get(path)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	if url.LongURL == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.Redirect(http.StatusFound, url.LongURL)
}

func (h *URLHandler) Create(c *gin.Context) {
	// TODO: answer here

	var url *entity.URL
	if err := c.ShouldBindJSON(&url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	urls, err := h.repo.Create(url.LongURL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": urls,
	})
}

func (h *URLHandler) CreateCustom(c *gin.Context) {
	// TODO: answer here

	var url *entity.URL
	if err := c.ShouldBindJSON(&url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	if url.LongURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": entity.ErrURLNotFound,
		})
	}
	urls, err := h.repo.CreateCustom(url.LongURL, url.ShortURL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	log.Println(urls)
	c.JSON(http.StatusOK, gin.H{
		"data": urls,
	})
}
