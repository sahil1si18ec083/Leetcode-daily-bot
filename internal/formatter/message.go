package formatter

import (
	"fmt"
	"regexp"
	"strings"

	"leetcode-daily-bot/internal/leetcode"
)

func Format(p leetcode.DailyProblem, aiText string) string {

	// 1. Remove markdown code fences
	clean := aiText
	clean = strings.ReplaceAll(clean, "```java", "")
	clean = strings.ReplaceAll(clean, "```", "")
	clean = strings.ReplaceAll(clean, "```cpp", "")
	clean = strings.ReplaceAll(clean, "```go", "")

	// 2. Remove HTML tags
	re := regexp.MustCompile(`<[^>]*>`)
	clean = re.ReplaceAllString(clean, "")

	// 3. Split explanation and code
	explanation := clean
	code := ""

	lines := strings.Split(clean, "\n")
	for i, line := range lines {
		if strings.Contains(line, "class ") || strings.Contains(line, "public") {
			explanation = strings.Join(lines[:i], "\n")
			code = strings.Join(lines[i:], "\n")
			break
		}
	}

	// 4. Trim explanation
	if len(explanation) > 500 {
		explanation = explanation[:500] + "..."
	}

	// 5. Build final SMS
	msg := "🔥 LeetCode Daily\n\n"
	msg += p.Title + " (" + p.Difficulty + ")\n\n"
	msg += "🧠 Idea:\n" + strings.TrimSpace(explanation) + "\n\n"
	msg += "💻 Java:\n" + strings.TrimSpace(code) + "\n\n"
	msg += "🔗 https://leetcode.com" + p.URL

	// 6. SMS safety limit
	if len(msg) > 60 {
		msg = msg[:60] + "\n\n(trimmed)"
	}

	fmt.Println(msg)

	return msg
}
