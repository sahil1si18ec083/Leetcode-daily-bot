package ai

type AiRequest struct {
	Contents []Content `json:"contents"`
}
type Content struct {
	Parts []Part `json:"parts"`
}
type Part struct {
	Text string `json:"text"`
}

type AiResponse struct {
	Candidates []Candidate `json:"candidates"`
}
type Candidate struct {
	Content Content `json:"content"`
}
