package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	// TODO: answer here
	router.GET("/movie", movieHandler)

	return router
}

// movie handler function
func movieHandler(c *gin.Context) {
	// ambil query string dari url
	movieGenre := c.Query("genre")
	movieCountry := c.Query("country")

	// kirim response ketika ada genre dan country
	if movieGenre != "" && movieCountry != "" {
		c.String(http.StatusOK, "Here result of query of movie with genre %s in country %s", movieGenre, movieCountry)
		return
	} else if movieGenre == "" && movieCountry != "" {
		// kirim response ketika hanya ada country saja
		c.String(http.StatusOK, "Here result of query of movie with genre general in country %s", movieCountry)
		return
	} else if movieGenre == "" && movieCountry == "" {
		c.String(http.StatusOK, "Here result of query of movie with genre general")
	} else {
		c.String(http.StatusOK, "Here result of query of movie with genre %s", movieGenre)
		return
	}
}

func main() {
	router := GetRouter()
	router.Run(":8080")
}
