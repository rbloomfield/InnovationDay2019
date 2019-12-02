package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nexmoinc/alice/models"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
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
				EventURL: []string{"https://en8fseqlqklpv.x.pipedream.net"},
				EndOnKey: "#",
			},
		}
		c.JSON(http.StatusOK, temp)
	})
	router.POST("/recoding", func(c *gin.Context) {
		// unmarshal the recoding message
		var json models.RecordingMessage
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusOK, nil)
			return
		}
	})

	router.Run(":8080")
}

// ngrok http 8080
