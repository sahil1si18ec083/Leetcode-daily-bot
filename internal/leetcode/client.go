package leetcode

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func FetchDaily() (DailyProblem, error) {
	url := "https://leetcode.com/graphql"
	query := "\n    query questionOfTodayV2 {\n  activeDailyCodingChallengeQuestion {\n    date\n    userStatus\n    link\n    question {\n      id: questionId\n      titleSlug\n      title\n      translatedTitle\n      questionFrontendId\n      paidOnly: isPaidOnly\n      difficulty\n      content\n      topicTags {\n        name\n        slug\n        nameTranslated: translatedName\n      }\n      status\n      isInMyFavorites: isFavor\n      acRate\n      frequency: freqBar\n    }\n  }\n}\n    "
	v := Variables{}
	requestPayload := DailyProblemRequest{
		Variables:     v,
		Query:         query,
		OperationName: "questionOfTodayV2",
	}
	body, err := json.Marshal(requestPayload)
	if err != nil {
		return DailyProblem{}, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Referer", "https://leetcode.com")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return DailyProblem{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return DailyProblem{}, fmt.Errorf("Error while fetching Leetcode Problem")
	}

	defer resp.Body.Close()
	d := DailyProblemResponse{}
	err = json.NewDecoder(resp.Body).Decode(&d)
	if err != nil {
		return DailyProblem{}, err
	}
	l := DailyProblem{}
	l.Difficulty = d.Data.ActiveDailyCodingChallengeQuestion.Question.Difficulty
	l.Title = d.Data.ActiveDailyCodingChallengeQuestion.Question.Title
	l.URL = d.Data.ActiveDailyCodingChallengeQuestion.Link
	l.ID = d.Data.ActiveDailyCodingChallengeQuestion.Question.ID
	l.Content = d.Data.ActiveDailyCodingChallengeQuestion.Question.Content
	return l, nil

}
