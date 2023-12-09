package scrape

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/Jiang-Gianni/lcs/db"
)

type Store struct {
	DB *sql.DB
	Q  *db.Queries
}

func NewStore() *Store {
	store, err := sql.Open("sqlite3", "store.db")
	if err != nil {
		log.Fatal(err)
	}

	return &Store{
		DB: store,
		Q:  db.New(store),
	}
}

func (s *Store) SaveQuestion(ctx context.Context, q Question, c QuestionContent,
	e QuestionEditor, h QuestionHints) error {

	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Transaction to store to db
	qt := s.Q.WithTx(tx)

	count, err := qt.CountQuestionByTitleSlug(ctx, q.TitleSlug)
	if err != nil {
		return err
	}
	if int(count) > 0 {
		return fmt.Errorf("%s has count %d stored", q.TitleSlug, count)
	}

	iqp := db.InsertQuestionParams{
		QuestionID: q.FrontendQuestionID,
		Link:       "https://leetcode.com/problems/" + q.TitleSlug,
		Title:      q.Title,
		TitleSlug:  q.TitleSlug,
		IsPaidOnly: q.PaidOnly,
		Difficulty: q.Difficulty,
		Content:    c.Content,
	}
	if err := qt.InsertQuestion(ctx, iqp); err != nil {
		return err
	}

	for _, cs := range e.CodeSnippets {
		iep := db.InsertEditorParams{
			QuestionID: q.FrontendQuestionID,
			Lang:       cs.Lang,
			LangSlug:   cs.LangSlug,
			Code:       cs.Code,
		}
		if err := qt.InsertEditor(ctx, iep); err != nil {
			return err
		}
	}

	for _, hint := range h.Hints {
		ihp := db.InsertHintParams{
			QuestionID: q.FrontendQuestionID,
			Hint:       hint,
		}
		if err := qt.InsertHint(ctx, ihp); err != nil {
			return err
		}
	}

	return tx.Commit()
}
