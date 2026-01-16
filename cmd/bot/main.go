package main

import (
	"leetcode-daily-bot/internal/ai"
	"leetcode-daily-bot/internal/formatter"
	"leetcode-daily-bot/internal/leetcode"
	"leetcode-daily-bot/internal/sms"
	"log"
	"os"

	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	var err error

	if err = godotenv.Load(); err != nil {

		fmt.Println("Error loading .env file")
	}
	res, err := leetcode.FetchDaily()
	if err != nil {
		log.Fatal("Error while generating Daily Problem")
	}

	aiResponse, err := ai.GenerateExplanation(res.Difficulty, res.ID, res.Title, res.URL, res.Content)

	message := formatter.Format(res, aiResponse)

	mode := os.Getenv("SMS_MODE")
	if mode == "" {
		log.Fatal("error while reading .env")
	}
	fmt.Println(mode)

	if mode == "fake" {
		err = sms.SendMessage(&sms.FakeSender{}, message)
	} else {
		err = sms.SendMessage(&sms.TwilioSender{}, message)
	}
	if err != nil {
		fmt.Println("error while sending message")
	}

}
