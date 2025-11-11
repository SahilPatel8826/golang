package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	token := os.Getenv("SLACK_BOT_TOKEN")
	channelID := os.Getenv("CHANNEL_ID")

	if token == "" || channelID == "" {
		fmt.Println("Missing SLACK_BOT_TOKEN or CHANNEL_ID environment variables.")
		return
	}

	api := slack.New(token)

	fileArr := []string{"sih.pdf"} // ✅ make sure correct path

	for _, filePath := range fileArr {
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Printf("❌ Failed to open %s: %v\n", filePath, err)
			continue
		}
		defer file.Close()

		info, _ := file.Stat()
		fmt.Printf("Uploading %s (%d bytes)\n", filePath, info.Size())
		if info.Size() == 0 {
			fmt.Printf("❌ Skipping %s: file size is 0 bytes\n", filePath)
			continue
		}

		params := slack.UploadFileV2Parameters{
			Channel:  channelID,
			Filename: info.Name(),
			Reader:   file,
		}

		uploadedFile, err := api.UploadFileV2(params)
		if err != nil {
			fmt.Printf("❌ Failed to upload %s: %v\n", filePath, err)
			continue
		}

		fmt.Printf("✅ Uploaded: ID=%s | Title=%s\n", uploadedFile.ID, uploadedFile.Title)
	}
}
