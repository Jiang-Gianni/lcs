package scrape

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Hints struct {
	Data struct {
		Question QuestionHints `json:"question"`
	} `json:"data"`
}

type QuestionHints struct {
	Hints []string `json:"hints"`
}

func GetQuestionHints(titleSlug string) (QuestionHints, error) {
	h := &Hints{}
	client := &http.Client{}
	var data = strings.NewReader(`{"query":"\n    query questionHints($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    hints\n  }\n}\n    ","variables":{"titleSlug":"` + titleSlug + `"},"operationName":"questionHints"}`)
	req, err := http.NewRequest("POST", "https://leetcode.com/graphql/", data)
	if err != nil {
		return h.Data.Question, err
	}
	req.Header.Set("content-type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return h.Data.Question, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(h)
	if err != nil {
		return h.Data.Question, err
	}
	return h.Data.Question, nil
}
