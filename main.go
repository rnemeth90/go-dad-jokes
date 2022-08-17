package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rnemeth90/go-dad-jokes"
)

func main() {
	router := gin.Default()
	router.GET("/", api.GetRandomJoke)
	router.GET("/random", api.GetRandomJoke)
	router.GET("/search:searchString", func(c *gin.Context) {
		s := c.Param("searchString")
		r := api.SearchJokes(s)
		c.String(http.StatusOK, r)
	})

	router.Run("localhost:8080")
}
