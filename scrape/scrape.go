package scrape

import (
	"context"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var timeout = time.Second

func Scrape() {
	s := NewStore()

	questions, err := GetQuestionList("0")
	if err != nil {
		log.Fatal(err)
	}

	for _, q := range questions {
		time.Sleep(timeout)

		fmt.Println("Working on ", q.TitleSlug)

		content, err := GetQuestionContent(q.TitleSlug)
		if err != nil {
			fmt.Println(q.TitleSlug, " content error: ", err)
			continue
		}

		editor, err := GetQuestionEditorData(q.TitleSlug)
		if err != nil {
			fmt.Println(q.TitleSlug, " editor error: ", err)
			continue
		}

		hints, err := GetQuestionHints(q.TitleSlug)
		if err != nil {
			fmt.Println(q.TitleSlug, " hints error: ", err)
			continue
		}

		if err := s.SaveQuestion(context.Background(), q, content, editor, hints); err != nil {
			fmt.Println(q.TitleSlug, " save error: ", err)
			continue
		}

	}
}
