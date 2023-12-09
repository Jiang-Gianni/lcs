package generate

import (
	"context"
	"regexp"

	"github.com/Jiang-Gianni/lcs/db"
	"github.com/Jiang-Gianni/lcs/scrape"
)

var funcNameRegexp = regexp.MustCompile(`func (.*?)\(`)

func Go() error {
	ctx := context.Background()
	s := scrape.NewStore()
	questions, err := s.Q.GetAllQuestion(ctx)
	if err != nil {
		return err
	}
	for _, q := range questions {
		editors, err := s.Q.GetEditors(ctx, db.GetEditorsParams{QuestionID: q.QuestionID, LangSlug: "golang"})
		if err != nil {
			return err
		}
		hints, err := s.Q.GetHints(ctx, q.QuestionID)
		if err != nil {
			return err
		}
		WriteFile("leetcode/go/"+q.TitleSlug+"_test.go", goTemplate("test", q, editors, hints))
	}
	return nil
}
