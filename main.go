package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	bot_token:= os.Getenv("SLACK_BOT_TOKEN")

	channelId := os.Getenv("SLACK_CHANNEL_ID")



	api := slack.New(bot_token)
	channelArr := []string{channelId}
	fileArr := []string{"dataset.pdf"}

	for i :=0; i<len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File: fileArr[i],
		}
		file, err := api.UploadFile(params)

		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}

		fmt.Printf("file - %s, URL:%s\n", file.Name, file.URL)

	}
}