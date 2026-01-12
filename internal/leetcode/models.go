package leetcode

type DailyProblem struct {
	Title      string `json:"title"`
	Difficulty string `json:"difficulty"`
	ID         string `json:"id"`
	URL        string `json:"url"`
	Content    string `json:"content"`
}
type DailyProblemResponse struct {
	Data Data `json:"data"`
}
type Data struct {
	ActiveDailyCodingChallengeQuestion ActiveDailyCodingChallengeQuestion `json:"activeDailyCodingChallengeQuestion"`
}
type ActiveDailyCodingChallengeQuestion struct {
	Link     string   `json:"link"`
	Question Question `json:"question"`
}
type Question struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Difficulty string `json:"difficulty"`
	Content    string `json:"content"`
}
type DailyProblemRequest struct {
	Variables     Variables `json:"variables"`
	OperationName string    `json:"operationName"`
	Query         string    `json:"query"`
}
type Variables struct {
}
