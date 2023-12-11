package scrape

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func Scrape(skip string) {
	defer func(start time.Time) {
		fmt.Println("Scraping and saving to db time: ", time.Since(start))
	}(time.Now())

	s := NewStore()

	questions, err := GetQuestionList(skip)
	if err != nil {
		log.Fatal(err)
	}

	wg := &sync.WaitGroup{}

	wg.Add(len(questions))

	for _, q := range questions {
		go s.GetQuestionData(q, wg)
	}

	wg.Wait()
}

func (s *Store) GetQuestionData(q Question, wg *sync.WaitGroup) {
	defer wg.Done()

	if q.PaidOnly {
		fmt.Printf("Skipping %s cause you got no money to pay for premium\n", q.TitleSlug)
		return
	}

	count, err := s.Q.CountQuestionByTitleSlug(context.Background(), q.TitleSlug)
	if err != nil {
		fmt.Println(q.TitleSlug, " count error: ", err)
		return
	}

	if int(count) > 0 {
		fmt.Println(q.TitleSlug, " is already stored")
		return
	}

	fmt.Println("Working on ", q.TitleSlug)

	var content QuestionContent

	var editor QuestionEditor

	var hints QuestionHints

	innerWg := &sync.WaitGroup{}
	innerWg.Add(3)

	go func() {
		defer innerWg.Done()

		content, err = GetQuestionContent(q.TitleSlug)
	}()

	go func() {
		defer innerWg.Done()

		editor, err = GetQuestionEditorData(q.TitleSlug)
	}()

	go func() {
		defer innerWg.Done()

		hints, err = GetQuestionHints(q.TitleSlug)
	}()

	innerWg.Wait()

	if err != nil {
		fmt.Println(q.TitleSlug, " content or editor or hints error: ", err)
		return
	}

	if err := s.SaveQuestion(context.Background(), q, content, editor, hints); err != nil {
		fmt.Println(q.TitleSlug, " save error: ", err)
		return
	}
}
