package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Movie struct {
	Title string
}

var movies = map[int]Movie{
	1: {
		"Spiderman",
	},
	2: {
		"Joker",
	},
	3: {
		"Escape Plan",
	},
	4: {
		"Lord of the Rings",
	},
}

var MovieListHandler = func(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": movies,
	}) // TODO: replace this
}

var MovieGetHandler = func(c *gin.Context) {
	// ambil id dari param
	id := c.Param("id")
	filmId, _ := strconv.Atoi(id)

	// jika ada fild dengan id yang sama
	if _, ok := movies[filmId]; ok {
		c.JSON(http.StatusOK, gin.H{
			"data": movies[filmId],
		}) // TODO: replace this
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "movie not found",
		})
		return
	}
}

func GetRouter() *gin.Engine {
	router := gin.Default()
	// TODO: answer here
	v1 := router.Group("/movie")
	{
		v1.GET("/list", MovieListHandler)
		// v1.GET("/get/:id", MovieGetHandler) //bisa cara ini
	}

	// bisa juga cara ini
	v2 := router.Group("/movie")
	{
		v2.GET("/get/:id", MovieGetHandler)
	}
	return router
}

func main() {
	router := GetRouter()

	router.Run(":8080")
}
