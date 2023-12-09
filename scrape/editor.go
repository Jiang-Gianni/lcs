package scrape

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Editor struct {
	Data struct {
		Question QuestionEditor `json:"question"`
	} `json:"data"`
}

type QuestionEditor struct {
	QuestionID         string `json:"questionId"`
	QuestionFrontendID string `json:"questionFrontendId"`
	CodeSnippets       []struct {
		Lang     string `json:"lang"`
		LangSlug string `json:"langSlug"`
		Code     string `json:"code"`
	} `json:"codeSnippets"`
	EnvInfo            string `json:"envInfo"`
	EnableRunCode      bool   `json:"enableRunCode"`
	HasFrontendPreview bool   `json:"hasFrontendPreview"`
	FrontendPreviews   string `json:"frontendPreviews"`
}

func GetQuestionEditorData(titleSlug string) (QuestionEditor, error) {
	ed := &Editor{}
	client := &http.Client{}
	var data = strings.NewReader(`{"query":"\n    query questionEditorData($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    questionId\n    questionFrontendId\n    codeSnippets {\n      lang\n      langSlug\n      code\n    }\n    envInfo\n    enableRunCode\n    hasFrontendPreview\n    frontendPreviews\n  }\n}\n    ","variables":{"titleSlug":"` + titleSlug + `"},"operationName":"questionEditorData"}`)
	req, err := http.NewRequest("POST", "https://leetcode.com/graphql/", data)
	if err != nil {
		return ed.Data.Question, err
	}
	req.Header.Set("content-type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return ed.Data.Question, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(ed)
	if err != nil {
		return ed.Data.Question, err
	}
	return ed.Data.Question, nil
}
