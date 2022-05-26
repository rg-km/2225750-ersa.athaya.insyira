package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Dari contoh yang telah diberikan, buatlah HTTP server sederhana dengan menggunakan Gin.
// Buatlah server yang memiliki address localhost dengan port 3000.
// Buatlah route "/hello" yang menampilkan JSON {"message": "hello"} dan "/world" yang menampilkan JSON {"message": "world"}

func GetGinRoute() *gin.Engine {

	// inisiasi gin
	r := gin.Default()

	// handler /hello
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello",
		})
	})

	// handler /world
	r.GET("/world", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "world",
		})
	})

	return r
}

func main() {
	r := GetGinRoute()
	r.Run("localhost:3000")
}
