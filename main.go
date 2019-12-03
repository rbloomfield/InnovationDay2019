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
const requestBin string = "https://en8fseqlqklpv.x.pipedream.net/"

var ConversationUUIDToRecordings map[string][]string = make(map[string][]string)

/*
func SendToAlice(SendToAlice []string) {

	body := struct {
		Name string `json:name`
		URL  string `json:url`
	}{
		Name: "",
	}
	req := http.NewRequest("POST", "url")
}
*/
func main() {

	//toAliceChannel := make(chan [3]string, 10)

	router := gin.Default()

	router.GET(eventURL, func(c *gin.Context) {

		speechAction := models.SpeechInput{
			Context:  []string{"name"},
			Language: "en-gb",
			// UUID:,
		}
		asrAction := models.NCCO{
			Action:       "input",
			Speech:       &speechAction,
			EndOnSilence: "2",
		}
		talk := models.NCCO{
			Action: "talk",
			Text:   "Please say your passphrase",
		}
		record := models.NCCO{
			Action:       "record",
			EventURL:     []string{"http://a9091a98.ngrok.io" + answerURL},
			EndOnKey:     "#",
			EndOnSilence: "2",
		}
		mainNCCO := []models.NCCO{
			models.NCCO{
				Action: "talk",
				Text:   "Please say your name",
			},
			asrAction,
			talk,
			record,
			talk,
			record,
			talk,
			record,
		}

		speechAction.UUID = []string{c.Query("uuid")}
		c.JSON(http.StatusOK,

			mainNCCO)
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
				// res := SendToAlice(i)
				// c.JSON(http.StatusOK, res)

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
