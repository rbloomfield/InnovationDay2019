package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nexmoinc/alice-client/models"
)

const port string = ":8080"
const eventURL string = "/event"
const answerURL string = "/answer"
const requestBin string = "https://en8fseqlqklpv.x.pipedream.net/"
const nameEventURL string = "/name"

type NameRecordings struct {
	name       string
	recordings []string
}

var ConversationUUIDToRecordings map[string]NameRecordings = make(map[string]NameRecordings)

func AddUser(i []string) bool {

	// body := struct {
	// 	Name string `json:name`
	// 	URL  string `json:url`
	// }{
	// 	Name: "",
	// }
	//	req := http.NewRequest("POST", "url")
	return false
}

func main() {

	//toAliceChannel := make(chan [3]string, 10)

	router := gin.Default()

	// return the NCCO to use
	router.GET(answerURL, func(c *gin.Context) {

		speechAction := models.SpeechInput{
			Context:  []string{"name"},
			Language: "en-gb",
			UUID:     []string{c.Query("uuid")},
		}
		asrAction := models.NCCO{
			Action:       "input",
			Speech:       &speechAction,
			EndOnSilence: "2",
			EventURL:     []string{"http://a9091a98.ngrok.io" + nameEventURL},
		}
		talk := models.NCCO{
			Action: "talk",
			Text:   "Please say your passphrase",
		}
		record := models.NCCO{
			Action:       "record",
			EventURL:     []string{"http://a9091a98.ngrok.io" + eventURL},
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

		//speechAction.UUID = []string{c.Query("uuid")}
		c.JSON(http.StatusOK,
			mainNCCO)
	})

	router.POST(nameEventURL, func(c *gin.Context) {
		var res models.ASRResponse
		if err := c.ShouldBindJSON(&res); err != nil {
			fmt.Println("\n\n WARNING bind error in " + nameEventURL)

		}
		fmt.Println("%#v", res)
		name := ""

		if len(res.Speech.Results) > 0 {
			name = res.Speech.Results[0].Text
		}
		ConversationUUIDToRecordings[res.UUID] = NameRecordings{name: name}
		c.JSON(http.StatusOK, nil)
	})

	// event url for Events
	router.POST(eventURL, func(c *gin.Context) {
		// unmarshal the recoding message
		var rec models.RecordingMessage
		if err := c.ShouldBindJSON(&rec); err != nil {
			c.JSON(http.StatusOK, nil)
			return
		}

		fmt.Printf("%#v", rec)
		if i, ok := ConversationUUIDToRecordings[rec.ConversationUUID]; ok {

			if i.recordings == nil {
				i.recordings = make([]string, 3)
			} else if len(i.recordings) > 0 {
				i.recordings = append(i.recordings, rec.RecordingURL)
			}
			if len(i.recordings) >= 3 {
				// here we should forward it on
				fmt.Println("Forward to alice")
				//res := AddUser(i)
				c.JSON(http.StatusOK, nil) //, res)

			}
			c.JSON(http.StatusOK, nil)
			return
		}
		// potential race issue
		//ConversationUUIDToRecordings[rec.ConversationUUID] = append(make([]string, 3), rec.RecordingURL)
		c.JSON(http.StatusOK, nil)

	})

	router.Run(port)
}

// ngrok http 8080
