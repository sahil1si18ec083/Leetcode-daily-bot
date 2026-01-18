package main

import (
	"leetcode-daily-bot/internal/ai"
	"leetcode-daily-bot/internal/formatter"
	"leetcode-daily-bot/internal/leetcode"
	"leetcode-daily-bot/internal/notify"
	"log"
	"os"
	"sync"

	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	errchan := make(chan error, 2)

	var err error
	var wg sync.WaitGroup
	wg.Add(2)

	if err = godotenv.Load(); err != nil {

		fmt.Println("Error loading .env file")
	}
	res, err := leetcode.FetchDaily()
	if err != nil {
		log.Fatal("Error while generating Daily Problem")
	}

	aiResponse, err := ai.GenerateExplanation(res.Difficulty, res.ID, res.Title, res.URL, res.Content)

	message := formatter.Format(res, aiResponse)

	sms_mode := os.Getenv("SMS_MODE")
	wapp_mode := os.Getenv("WAPP_MODE")
	if sms_mode == "" || wapp_mode == "" {
		fmt.Println("error while reading .env")
	}
	fmt.Println(sms_mode)
	go func() {
		defer wg.Done()
		if sms_mode == "fake" {
			notify.SendMessage(&notify.FakeSender{}, message, errchan)
		} else {
			notify.SendMessage(&notify.TwilioSender{}, message, errchan)
		}
	}()
	go func() {
		defer wg.Done()
		if wapp_mode == "fake" {
			notify.SendMessage(&notify.FakeSender{}, message, errchan)
		} else {
			notify.SendMessage(&notify.WhatsAppSender{}, message, errchan)
		}

	}()
	wg.Wait()
	close(errchan)
	for err := range errchan {
		if err != nil {
			fmt.Println(err)
			fmt.Println("error while sending message")
		}
	}

}
