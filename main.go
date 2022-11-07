package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"lotto/slack"
	"lotto/util"
	"net/http"
	"sort"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.GET("/lotto", func(c *gin.Context) {
		luckyNumbers := sendLottoNumber(false)
		c.JSON(http.StatusOK, gin.H{"data": luckyNumbers})
	})

	r.POST("/lotto", func(c *gin.Context) {
		sendSlack := c.DefaultQuery("sendSlack", "false")
		luckyNumbers := sendLottoNumber(util.ParseBool(sendSlack))
		c.JSON(http.StatusOK, gin.H{"data": luckyNumbers})
	})

	r.POST("/slack/webhook/:name", func(c *gin.Context) {
		webhookName := c.Param("name")
		// TODO: register webhook with name
		c.JSON(http.StatusOK, gin.H{"data": webhookName})
	})

	r.Run()
}

func sendLottoNumber(sendSlack bool) string {
	luckyNumbers := util.PickRandomNumbers(1, 45, 6)
	sort.Ints(luckyNumbers)
	luckyNumbersString := util.IntArrayToString(luckyNumbers, " ")
	if sendSlack {
		slack.SendMessage(fmt.Sprintf("Your lucky numbers are `%s`", luckyNumbersString))
	}
	return luckyNumbersString
}
