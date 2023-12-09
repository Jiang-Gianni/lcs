package scrape

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Title struct {
	Data struct {
		Question QuestionTitle `json:"question"`
	} `json:"data"`
}

type QuestionTitle struct {
	QuestionID         string `json:"questionId"`
	QuestionFrontendID string `json:"questionFrontendId"`
	Title              string `json:"title"`
	TitleSlug          string `json:"titleSlug"`
	IsPaidOnly         bool   `json:"isPaidOnly"`
	Difficulty         string `json:"difficulty"`
	Likes              int    `json:"likes"`
	Dislikes           int    `json:"dislikes"`
	CategoryTitle      string `json:"categoryTitle"`
}

func GetQuestionTitle(titleSlug string) (QuestionTitle, error) {
	t := &Title{}
	client := &http.Client{}
	var data = strings.NewReader(`{"query":"\n    query questionTitle($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    questionId\n    questionFrontendId\n    title\n    titleSlug\n    isPaidOnly\n    difficulty\n    likes\n    dislikes\n    categoryTitle\n  }\n}\n    ","variables":{"titleSlug":"` + titleSlug + `"},"operationName":"questionTitle"}`)
	req, err := http.NewRequest("POST", "https://leetcode.com/graphql/", data)
	if err != nil {
		return t.Data.Question, err
	}
	req.Header.Set("content-type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return t.Data.Question, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(t)
	if err != nil {
		return t.Data.Question, err
	}
	return t.Data.Question, nil
}
