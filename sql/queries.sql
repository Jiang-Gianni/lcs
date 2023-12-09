-- name: GetAllQuestion :many
select * from question;

-- name: CountQuestionByTitleSlug :one
select count(*) from question where title_slug = ?;

-- name: InsertQuestion :exec
insert into question(question_id, link, title, title_slug, is_paid_only, difficulty, content) values (?, ?, ?, ?, ?, ?, ?);

-- name: InsertHint :exec
insert into hint(question_id, hint) values (?, ?);

-- name: InsertEditor :exec
insert into editor(question_id, lang, lang_slug, code) values (?, ?, ?, ?);

-- name: GetHints :many
select * from hint where question_id = ?;

-- name: GetEditors :many
select * from editor where question_id = ? and lang_slug = ?;