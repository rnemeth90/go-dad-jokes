package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"sort"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", getRandomJoke)
	router.GET("/random", getRandomJoke)
	router.GET("/search:searchString", func(c *gin.Context) {
		s := c.Param("searchString")
		r := searchJokes(s)
		c.String(http.StatusOK, r)
	})

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
		log.Fatal("Error decoding JSON: ", err)
	}
	n := rand.Int() % len(jokes)
	c.IndentedJSON(http.StatusOK, jokes[n])
}

func searchJokes(searchString string) string {

	log.SetPrefix("searchJokes(): ")
	jokes := make([]string, 0)
	// results := make([]string, 0)
	var result string

	jsonFile, err := ioutil.ReadFile("./jokes.json")
	if err != nil {
		log.Fatal("Error opening JSON: ", err)
	}

	err = json.Unmarshal(jsonFile, &jokes)
	if err != nil {
		log.Fatal("Error decoding JSON: ", err)
	}
	sort.Strings(jokes)

	x := sort.SearchStrings(jokes, searchString)
	if x > 0 {
		result = jokes[x]
	}

	return result
}
