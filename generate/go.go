package generate

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/Jiang-Gianni/lcs/db"
	"github.com/Jiang-Gianni/lcs/scrape"
)

const leetcodego = "leetcode/go/"

func Go(from, to int) error {
	if err := os.MkdirAll(leetcodego, 0777); err != nil {
		return err
	}
	ctx := context.Background()
	s := scrape.NewStore()

	questions, err := s.Q.GetQuestions(ctx, db.GetQuestionsParams{From: int64(from), To: int64(to)})
	if err != nil {
		return err
	}

	for _, q := range questions {
		fmt.Println("working on ", q.TitleSlug)
		editors, err := s.Q.GetEditors(ctx, db.GetEditorsParams{QuestionID: q.QuestionID, LangSlug: "golang"})
		if len(editors) == 0 {
			fmt.Printf("No editor for golang and question %s\n", q.TitleSlug)
			continue
		}
		if err != nil {
			return err
		}
		hints, err := s.Q.GetHints(ctx, q.QuestionID)
		if err != nil {
			return err
		}
		WriteFile(
			leetcodego+PadWithZero(q.QuestionID, 4)+"_"+q.TitleSlug+"_test.go", goTemplate(q, editors, hints),
		)
	}
	return nil
}

func PadWithZero(value string, n int) string {
	zeroes := n - len(value)
	if zeroes < 0 {
		zeroes = 0
	}
	return strings.Repeat("0", zeroes) + value
}
