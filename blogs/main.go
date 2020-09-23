package main

import (
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
	"log"
	"encoding/json"
)

var port = ":8080"

func main() {
	r:= gin.Default()

	r.POST("/blogs", clientServer)

	r.Run(port)
}

type clientInput struct {
	Author 		string `json:"author"`
	Title 		string `json:"title"`
	Comments 	[]struct{
		Message string `json:"message"`
	} `json:"comments"`
}

func clientServer(c *gin.Context) {
	var input clientInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
	}

	authorLog, err := os.OpenFile("author.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	titleLog, err := os.OpenFile("title.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	messageLog, err := os.OpenFile("message.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer authorLog.Close()
	defer titleLog.Close()
	defer messageLog.Close()

	var fullPath = string(c.FullPath())
	
	var authorStringLog = `Success: POST http://127.0.0.0`+ port + fullPath +` {"author" : "` + input.Author + `"}`
	var titleStringLog = `Success: POST http://127.0.0.0`+ port + fullPath +` {"title" : "`+ input.Title +`"}`
	m, _ := json.Marshal(input.Comments)
	var messageOutput = string(m)
	var messageStringLog = `Success: POST http://127.0.0.0`+ port + fullPath +` {"comments" : `+ messageOutput +`}`
	
	log.SetOutput(authorLog)
	log.Print(authorStringLog)
	log.SetOutput(titleLog)
	log.Print(titleStringLog)
	log.SetOutput(messageLog)
	log.Print(messageStringLog)

}