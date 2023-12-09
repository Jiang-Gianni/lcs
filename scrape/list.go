package scrape

import (
	"encoding/json"
	"net/http"
	"strings"
)

type List struct {
	Data struct {
		ProblemsetQuestionList struct {
			Total     int        `json:"total"`
			Questions []Question `json:"questions"`
		} `json:"problemsetQuestionList"`
	} `json:"data"`
}

type Question struct {
	AcRate             float64     `json:"acRate"`
	Difficulty         string      `json:"difficulty"`
	FreqBar            interface{} `json:"freqBar"`
	FrontendQuestionID string      `json:"frontendQuestionId"`
	IsFavor            bool        `json:"isFavor"`
	PaidOnly           bool        `json:"paidOnly"`
	Status             string      `json:"status"`
	Title              string      `json:"title"`
	TitleSlug          string      `json:"titleSlug"`
	TopicTags          []struct {
		Name string `json:"name"`
		ID   string `json:"id"`
		Slug string `json:"slug"`
	} `json:"topicTags"`
	HasSolution      bool `json:"hasSolution"`
	HasVideoSolution bool `json:"hasVideoSolution"`
}

func GetQuestionList(skip string) ([]Question, error) {
	l := &List{}
	client := &http.Client{}
	var data = strings.NewReader(`{"query":"\n    query problemsetQuestionList($categorySlug: String, $limit: Int, $skip: Int, $filters: QuestionListFilterInput) {\n  problemsetQuestionList: questionList(\n    categorySlug: $categorySlug\n    limit: $limit\n    skip: $skip\n    filters: $filters\n  ) {\n    total: totalNum\n    questions: data {\n      acRate\n      difficulty\n      freqBar\n      frontendQuestionId: questionFrontendId\n      isFavor\n      paidOnly: isPaidOnly\n      status\n      title\n      titleSlug\n      topicTags {\n        name\n        id\n        slug\n      }\n      hasSolution\n      hasVideoSolution\n    }\n  }\n}\n    ","variables":{"categorySlug":"all-code-essentials","skip":` + skip + `,"limit":50,"filters":{}},"operationName":"problemsetQuestionList"}`)
	req, err := http.NewRequest("POST", "https://leetcode.com/graphql/", data)
	if err != nil {
		return nil, err
	}
	req.Header.Set("content-type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(l)
	if err != nil {
		return nil, err
	}
	return l.Data.ProblemsetQuestionList.Questions, nil
}
