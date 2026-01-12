package main

import (
	"leetcode-daily-bot/internal/sms"
	"log"
	"os"

	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	var err error
	err = godotenv.Load()
	if err != nil {
		fmt.Println(err)
		log.Fatal("Error loading .env file")
	}
	message := "test message"

	mode := os.Getenv("SMS_MODE")
	if mode == "" {
		log.Fatal("error while reading .env")
	}

	if mode == "fake" {
		err = sms.SendMessage(&sms.FakeSender{}, message)
	} else {
		err = sms.SendMessage(&sms.TwilioSender{}, message)
	}
	if err != nil {
		fmt.Println("error while sending message")
	}

}
