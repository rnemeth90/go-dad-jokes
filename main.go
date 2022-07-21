package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", getRandomJoke)
	router.Run("localhost:8080")
}

func getRandomJoke(c *gin.Context) {
	rand.Seed(time.Now().Unix())
	log.SetPrefix("getRandomJoke(): ")

	jsonFile, err := ioutil.ReadFile("./jokes.json")
	if err != nil {
		log.Fatal("Error opening the file: ", err)
	}

	jokes := make([]string, 0)
	err = json.Unmarshal(jsonFile, &jokes)
	if err != nil {
		log.Fatal("Error decoding json: ", err)
	}
	n := rand.Int() % len(jokes)
	c.IndentedJSON(http.StatusOK, jokes[n])
}
