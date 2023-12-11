package scrape

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Content struct {
	Data struct {
		Question QuestionContent `json:"question"`
	} `json:"data"`
}

type QuestionContent struct {
	Content      string        `json:"content"`
	MysqlSchemas []interface{} `json:"mysqlSchemas"`
	DataSchemas  []interface{} `json:"dataSchemas"`
}

func GetQuestionContent(titleSlug string) (QuestionContent, error) {
	c := &Content{}
	client := &http.Client{}

	var data = strings.NewReader(`{"query":"\n    query questionContent($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    content\n    mysqlSchemas\n    dataSchemas\n  }\n}\n    ","variables":{"titleSlug":"` + titleSlug + `"},"operationName":"questionContent"}`)

	req, err := http.NewRequest("POST", "https://leetcode.com/graphql/", data)
	if err != nil {
		return c.Data.Question, err
	}

	req.Header.Set("content-type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return c.Data.Question, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(c)
	if err != nil {
		return c.Data.Question, err
	}

	return c.Data.Question, err
}
