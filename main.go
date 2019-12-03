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

var ConversationUUIDToRecordings map[string][]string = make(map[string][]string)

func main() {

	toAliceChannel := make(chan [3]string, 10)

	router := gin.Default()

	talk := models.NNCO{
		Action: "talk",
		Text:   "Please say your passphrase",
	}
	record := models.NNCO{
		Action:       "record",
		EventURL:     []string{"http://a9091a98.ngrok.io" + answerURL},
		EndOnKey:     "#",
		EndOnSilence: "2",
	}
	mainNCCO := []models.NNCO{
		talk,
		record,
		talk,
		record,
		talk,
		record,
	}
	router.GET(eventURL, func(c *gin.Context) {
		c.JSON(http.StatusOK, mainNCCO)
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
		if i, ok := ConversationUUIDToRecordings[json.ConversationUUID]; ok {
			ConversationUUIDToRecordings[json.ConversationUUID] = append(i, json.RecordingURL)
			if len(i) >= 3 {
				// here we should forward it on
				fmt.Println("Forward to alice")

			}
			c.JSON(http.StatusOK, nil)
			return
		}
		ConversationUUIDToRecordings[json.ConversationUUID] = append(make([]string, 3), json.RecordingURL)
		c.JSON(http.StatusOK, nil)

	})

	router.Run(port)
}

// ngrok http 8080
