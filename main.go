package main

import "github.com/gin-gonic/gin"

import "net/http"

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

	router.Run(":8080")
}
