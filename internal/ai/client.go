package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func GenerateExplanation(difficulty string, id string, title string, url string, content string) (string, error) {

	model := "gemini-2.5-flash"
	key := os.Getenv("GEMINI_API_KEY")
	if key == "" {
		return "", fmt.Errorf("Empty GEMINI_API_KEY")
	}
	gemini_url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/%s:generateContent?key=%s", model, key)
	prompt := fmt.Sprintf(`
You are a LeetCode tutor.

Give a VERY SHORT explanation and a concise Java solution.
The entire response MUST be under 700 characters.

Format strictly like this:

Idea:
<2-3 lines explanation>

Java:
<short Java solution>

Problem:
Title: %s
Difficulty: %s
Link: %s

Description:
%s
`, title, difficulty, url, content)

	p := Part{Text: prompt}
	c := Content{Parts: []Part{p}}
	requestPayload := AiRequest{
		Contents: []Content{c},
	}
	body, err := json.Marshal(requestPayload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, gemini_url, bytes.NewBuffer(body))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Error while fetching Leetcode Problem")
	}

	defer resp.Body.Close()
	r := AiResponse{}
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return "", err
	}

	return r.Candidates[0].Content.Parts[0].Text, nil

}
