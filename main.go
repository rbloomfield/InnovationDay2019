package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty"
	"github.com/nexmoinc/alice-client/models"
)

const port string = ":8080"
const NgrokAddress string = "http://a9091a98.ngrok.io"
const EventURL string = "/event"
const AnswerURL string = "/answer"
const RequestBin string = "https://en8fseqlqklpv.x.pipedream.net/"
const NameEventURL string = "/name"
const AliceURL string = "http://ec2-54-165-140-92.compute-1.amazonaws.com:8080"

type NameRecordings struct {
	Name       string   `json:"name"`
	Recordings []string `json:"urls"`
}

var ConversationUUIDToRecordings map[string]*NameRecordings = make(map[string]*NameRecordings)

func AddUser(i NameRecordings) bool {

	// body := struct {
	// 	Name string `json:name`
	// 	URL  string `json:url`
	// }{
	// 	Name: "",
	// }
	// req := http.NewRequest("POST", "url")
	// Create a Resty Client
	client := resty.New()

	res, err := client.R().
		EnableTrace().
		SetBody(i).
		Post(AliceURL + "/names")
	if err != nil {
		fmt.Println("error : ", err.Error())
	}
	return res.StatusCode() == http.StatusOK
}

func main() {

	//toAliceChannel := make(chan [3]string, 10)

	router := gin.Default()

	// return the NCCO to use
	router.GET(AnswerURL, func(c *gin.Context) {

		speechAction := models.SpeechInput{
			Context:  []string{"name"},
			Language: "en-gb",
			UUID:     []string{c.Query("uuid")},
		}
		asrAction := models.NCCO{
			Action:       "input",
			Speech:       &speechAction,
			EndOnSilence: "2",
			EventURL:     []string{"http://a9091a98.ngrok.io" + NameEventURL},
		}
		talk := models.NCCO{
			Action: "talk",
			Text:   "Please say your passphrase",
		}
		record := models.NCCO{
			Action:       "record",
			EventURL:     []string{"http://a9091a98.ngrok.io" + EventURL},
			EndOnKey:     "#",
			EndOnSilence: "2",
			Format:       "wav",
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

	router.POST(NameEventURL, func(c *gin.Context) {
		var res models.ASRResponse
		if err := c.ShouldBindJSON(&res); err != nil {
			fmt.Println("\n\n WARNING bind error in " + NameEventURL)

		}
		// fmt.Println("[DEBUG] %#v", res)
		name := "name not found"
		if len(res.Speech.Results) > 0 {
			name = res.Speech.Results[0].Text
		}

		fmt.Printf("Creating entry for : %s, name: %s \n", res.ConversationUUID, name)
		ConversationUUIDToRecordings[res.ConversationUUID] = &NameRecordings{Name: name}
		c.JSON(http.StatusOK, nil)
	})

	// event url for Events
	router.POST(EventURL, func(c *gin.Context) {

		// unmarshal the recoding message
		var rec models.RecordingMessage
		if err := c.ShouldBindJSON(&rec); err != nil {
			// bad request
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		// look up if this user has been added to our cache
		if nameRecordings, ok := ConversationUUIDToRecordings[rec.ConversationUUID]; ok {

			// add the api_key and api_secret for Nexmo
			recordingURL := rec.RecordingURL + "?api_key=4e2ebfb7&api_secret=ctb5LNo8cYOcf82k"
			// fmt.Printf("[DEBUG] appending to recordings for: %s\n", rec.ConversationUUID)
			nameRecordings.Recordings = append(nameRecordings.Recordings, recordingURL)

			if len(nameRecordings.Recordings) >= 3 { // enough instances to pass on
				// fmt.Println("[DEBUG] Forward to alice")
				res := AddUser(*nameRecordings)
				c.JSON(http.StatusOK, res)
			}
			// fmt.Printf("[DEBUG]length : %v ", len(nameRecordings.Recordings))
			c.JSON(http.StatusOK, nil)
			return
		}

		fmt.Println("[Eror] unrecognized converation uuid in map")
		c.JSON(http.StatusBadRequest, nil)
	})

	router.Run(port)
}

// ngrok http 8080
