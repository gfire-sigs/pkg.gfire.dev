// Code generated by qtc from "index.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/index.qtpl:1
package views

//line views/index.qtpl:1
import "github.com/gfire-sigs/pkg.gfire.dev/types"

//line views/index.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/index.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/index.qtpl:4
type Module = types.Module

//line views/index.qtpl:7
func StreamIndex(qw422016 *qt422016.Writer, Modules []*Module) {
//line views/index.qtpl:7
	qw422016.N().S(`<!DOCTYPE html>
`)
//line views/index.qtpl:8
	qw422016.N().S(`<html lang="en"><head><meta charset="UTF-8"><meta http-equiv="X-UA-Compatible" content="IE=edge"><meta name="viewport" content="width=device-width, initial-scale=1.0"><title>pkg.gfire.dev Module Index</title><meta name="og:title" content="pkg.gfire.dev Module Index"><meta name="og:description" content="A Module index for Go Modules hosted on pkg.gfire.dev"><meta name="description" content="A Module index for Go Modules hosted on pkg.gfire.dev"><link rel="stylesheet" href="/_static/main.css"></head><body><h1>pkg.gfire.dev Module Index</h1><span class="description">A Module index for Go Modules hosted on pkg.gfire.dev</span><span class="description">Index Repository: <a href="https://github.com/gfire-sigs/pkg.gfire.dev">https://github.com/gfire-sigs/pkg.gfire.dev</a></span><br/><div class="box"><ul>`)
//line views/index.qtpl:27
	for _, p := range Modules {
//line views/index.qtpl:27
		qw422016.N().S(`<li><a href="/`)
//line views/index.qtpl:29
		qw422016.E().S(p.Path)
//line views/index.qtpl:29
		qw422016.N().S(`">`)
//line views/index.qtpl:29
		qw422016.E().S(p.Root)
//line views/index.qtpl:29
		qw422016.N().S(`</a></li>`)
//line views/index.qtpl:31
	}
//line views/index.qtpl:31
	qw422016.N().S(`</ul></div></body></html>`)
//line views/index.qtpl:36
	qw422016.N().S(`
`)
//line views/index.qtpl:37
}

//line views/index.qtpl:37
func WriteIndex(qq422016 qtio422016.Writer, Modules []*Module) {
//line views/index.qtpl:37
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/index.qtpl:37
	StreamIndex(qw422016, Modules)
//line views/index.qtpl:37
	qt422016.ReleaseWriter(qw422016)
//line views/index.qtpl:37
}

//line views/index.qtpl:37
func Index(Modules []*Module) string {
//line views/index.qtpl:37
	qb422016 := qt422016.AcquireByteBuffer()
//line views/index.qtpl:37
	WriteIndex(qb422016, Modules)
//line views/index.qtpl:37
	qs422016 := string(qb422016.B)
//line views/index.qtpl:37
	qt422016.ReleaseByteBuffer(qb422016)
//line views/index.qtpl:37
	return qs422016
//line views/index.qtpl:37
}
