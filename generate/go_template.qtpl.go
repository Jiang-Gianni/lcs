// Code generated by qtc from "go_template.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line generate/go_template.qtpl:1
package generate

//line generate/go_template.qtpl:1
import "github.com/Jiang-Gianni/lcs/db"

//line generate/go_template.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line generate/go_template.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line generate/go_template.qtpl:3
func streamgoTemplate(qw422016 *qt422016.Writer, q db.Question, editors []db.Editor, hints []db.Hint) {
//line generate/go_template.qtpl:3
	qw422016.N().S(`//go:build exclude

/*
`)
//line generate/go_template.qtpl:6
	qw422016.E().S(q.QuestionID)
//line generate/go_template.qtpl:6
	qw422016.N().S(` - `)
//line generate/go_template.qtpl:6
	qw422016.E().S(q.Title)
//line generate/go_template.qtpl:6
	qw422016.N().S(` (`)
//line generate/go_template.qtpl:6
	qw422016.E().S(q.Difficulty)
//line generate/go_template.qtpl:6
	qw422016.N().S(`)
`)
//line generate/go_template.qtpl:7
	qw422016.E().S(q.TitleSlug)
//line generate/go_template.qtpl:7
	qw422016.N().S(`
`)
//line generate/go_template.qtpl:8
	qw422016.E().S(q.Link)
//line generate/go_template.qtpl:8
	qw422016.N().S(`
`)
//line generate/go_template.qtpl:9
	for i, h := range hints {
//line generate/go_template.qtpl:9
		qw422016.N().S(`
Hint `)
//line generate/go_template.qtpl:10
		qw422016.E().V(i)
//line generate/go_template.qtpl:10
		qw422016.N().S(`
`)
//line generate/go_template.qtpl:11
		qw422016.E().S(h.Hint)
//line generate/go_template.qtpl:11
		qw422016.N().S(`
`)
//line generate/go_template.qtpl:12
	}
//line generate/go_template.qtpl:12
	qw422016.N().S(`

*/
package solution

import "testing"

`)
//line generate/go_template.qtpl:19
	for _, e := range editors {
//line generate/go_template.qtpl:19
		qw422016.N().S(`
`)
//line generate/go_template.qtpl:21
		gs := GetGoSignature(e.Code)

//line generate/go_template.qtpl:22
		qw422016.N().S(`
func Test`)
//line generate/go_template.qtpl:23
		qw422016.E().S(gs.upperName)
//line generate/go_template.qtpl:23
		qw422016.N().S(`(t *testing.T) {
	testCases := []struct {
		name string`)
//line generate/go_template.qtpl:25
		for _, p := range gs.parameters {
//line generate/go_template.qtpl:25
			qw422016.N().S(`
        `)
//line generate/go_template.qtpl:26
			qw422016.E().S(p)
//line generate/go_template.qtpl:26
		}
//line generate/go_template.qtpl:26
		qw422016.N().S(`
        expected `)
//line generate/go_template.qtpl:27
		qw422016.E().S(gs.out)
//line generate/go_template.qtpl:27
		qw422016.N().S(`
	}{
		{
            name: "First",
        },
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
		})
	}

}

`)
//line generate/go_template.qtpl:41
		qw422016.E().S(e.Code)
//line generate/go_template.qtpl:41
		qw422016.N().S(`
`)
//line generate/go_template.qtpl:42
	}
//line generate/go_template.qtpl:42
}

//line generate/go_template.qtpl:42
func writegoTemplate(qq422016 qtio422016.Writer, q db.Question, editors []db.Editor, hints []db.Hint) {
//line generate/go_template.qtpl:42
	qw422016 := qt422016.AcquireWriter(qq422016)
//line generate/go_template.qtpl:42
	streamgoTemplate(qw422016, q, editors, hints)
//line generate/go_template.qtpl:42
	qt422016.ReleaseWriter(qw422016)
//line generate/go_template.qtpl:42
}

//line generate/go_template.qtpl:42
func goTemplate(q db.Question, editors []db.Editor, hints []db.Hint) string {
//line generate/go_template.qtpl:42
	qb422016 := qt422016.AcquireByteBuffer()
//line generate/go_template.qtpl:42
	writegoTemplate(qb422016, q, editors, hints)
//line generate/go_template.qtpl:42
	qs422016 := string(qb422016.B)
//line generate/go_template.qtpl:42
	qt422016.ReleaseByteBuffer(qb422016)
//line generate/go_template.qtpl:42
	return qs422016
//line generate/go_template.qtpl:42
}
