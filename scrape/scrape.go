package scrape

import (
	"context"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var timeout = time.Millisecond * 250

func Scrape(skip string) {
	s := NewStore()

	questions, err := GetQuestionList(skip)
	if err != nil {
		log.Fatal(err)
	}

	for _, q := range questions {
		if q.PaidOnly {
			fmt.Printf("Skipping %s cause you got no money to pay for premium\n", q.TitleSlug)
			continue
		}

		count, err := s.Q.CountQuestionByTitleSlug(context.Background(), q.TitleSlug)
		if err != nil {
			fmt.Println(q.TitleSlug, " count error: ", err)
			continue
		}
		if int(count) > 0 {
			fmt.Println(q.TitleSlug, " is already stored")
			continue
		}

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
