package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var RequestMethodHandler = func(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": c.Request.Method,
	})
}

func GetGinRoute() *gin.Engine {
	r := gin.Default()

	// GET
	r.GET("/get", getRequest)

	// POST
	r.POST("/post", postRequest)

	// PUT
	r.PUT("/put", putRequest)

	// DELETE
	r.DELETE("/delete", deleteRequest)

	// PATCH
	r.PATCH("/patch", patchRequest)

	// HEAD
	r.HEAD("/head", headRequest)

	// OPTIONS
	r.OPTIONS("/options", optionsRequest)

	return r
}

// hanlder /get
func getRequest(c *gin.Context) {
	// bikin response
	c.JSON(http.StatusOK, gin.H{
		"message": "GET",
	})
}

// hanlder /post
func postRequest(c *gin.Context) {
	// bikin response
	c.JSON(http.StatusOK, gin.H{
		"message": "POST",
	})
}

// handler /put
func putRequest(c *gin.Context) {
	// bikin response
	c.JSON(http.StatusOK, gin.H{
		"message": "PUT",
	})
}

// handler /delete
func deleteRequest(c *gin.Context) {
	// bikin response
	c.JSON(http.StatusOK, gin.H{
		"message": "DELETE",
	})
}

// handler /patch
func patchRequest(c *gin.Context) {
	// bikin response
	c.JSON(http.StatusOK, gin.H{
		"message": "PATCH",
	})
}

// handler /head
func headRequest(c *gin.Context) {
	// bikin response
	c.JSON(http.StatusOK, gin.H{
		"message": "HEAD",
	})
}

// handler /options
func optionsRequest(c *gin.Context) {
	// bikin response
	c.JSON(http.StatusOK, gin.H{
		"message": "OPTIONS",
	})
}

func main() {
	r := GetGinRoute()
	r.Run("localhost:3000")
}
