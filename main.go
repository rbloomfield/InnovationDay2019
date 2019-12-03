package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nexmoinc/alice-client/models"
)

const port string = ":8080"
const eventURL string = "/"
const answerURL string = "/answer"

func main() {
	router := gin.Default()

	router.GET(eventURL, func(c *gin.Context) {
		temp := []struct {
			Text     string   `json:"text,omitempty"`
			Action   string   `json:"action"`
			EventURL []string `json:"eventUrl,omitempty"`
			EndOnKey string   `json:"endOnKey,omitempty"`
		}{
			{
				Action: "talk",
				Text:   "Please say your passphrase",
			},
			{
				Action:   "record",
				EventURL: []string{"http://a9091a98.ngrok.io" + answerURL},
				EndOnKey: "#",
			},
		}
		c.JSON(http.StatusOK, temp)
	})

	router.POST(answerURL, func(c *gin.Context) {

		fmt.Println("Answer URL hit")
		// unmarshal the recoding message
		var json models.RecordingMessage
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusOK, nil)
			return
		}
		fmt.Printf("%#v", json)
	})

	router.Run(port)
}

// ngrok http 8080
