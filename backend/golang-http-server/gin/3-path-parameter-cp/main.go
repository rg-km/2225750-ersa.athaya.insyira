package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name    string
	Country string
	Age     int
}

var data = map[int]User{
	1: {
		Name:    "Roger",
		Country: "United States",
		Age:     70,
	},
	2: {
		Name:    "Tony",
		Country: "United States",
		Age:     40,
	},
	3: {
		Name:    "Asri",
		Country: "Indonesia",
		Age:     30,
	},
}

func ProfileHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		// amvil id dari url
		myParam := c.Param("id")
		// convert string ke int
		myId, _ := strconv.Atoi(myParam)

		// jika id tidak ditemukan
		if _, ok := data[myId]; !ok {
			c.String(http.StatusNotFound, "data not found")
			return
		} else {
			c.String(http.StatusOK, "Name: %s, Country: %s, Age: %d", data[myId].Name, data[myId].Country, data[myId].Age)
		}
	}
}

func GetRouter() *gin.Engine {
	r := gin.Default()

	// get ke /profile
	r.GET("/profile/:id", ProfileHandler())

	return r
}

func main() {
	router := GetRouter()
	router.Run(":8080")
}
