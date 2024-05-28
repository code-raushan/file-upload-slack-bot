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
	fileArr := []string{"dataset.pdf"}

	for _, file := range fileArr {
		fileData, err := os.Open(file)
		fileInfo, _ := fileData.Stat()
		params := slack.UploadFileV2Parameters{
			Channel: channelId,
			Filename: fileData.Name(),
			FileSize: int(fileInfo.Size()),
			File: file,
		}

		if err != nil {
			fmt.Printf("Error while opening file: %v\n", err)
			continue
		}
		defer fileData.Close()

		uploadRes, err := api.UploadFileV2(params)

		if err != nil {
			fmt.Printf("Error uploading file: %v\n", err)
			continue
		}
		
		fmt.Printf("File - %s uploaded", uploadRes.Title)
	}
}