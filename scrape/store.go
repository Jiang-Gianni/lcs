package scrape

import (
	"context"
	"database/sql"
	"log"

	"github.com/Jiang-Gianni/lcs/db"
)

type Store struct {
	DB *sql.DB
	Q  *db.Queries
}

const initTables = `create table if not exists question(
    id integer primary key not null,
    question_id text not null,
    link text not null,
    title text not null,
    title_slug text not null,
    is_paid_only boolean not null,
    difficulty text not null,
    content text not null
);

create table if not exists hint(
    id integer primary key not null,
    question_id text not null,
    hint text not null
);

create table if not exists editor(
    id integer primary key not null,
    question_id text not null,
    lang text not null,
    lang_slug text not null,
    code text not null
);
`

func NewStore() *Store {
	store, err := sql.Open("sqlite3", "store.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = store.Exec(initTables)
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
