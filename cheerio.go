//go:generate go-bindata -debug -pkg cheerio -o bindata.go ./dist
package cheerio

import (
	"gopkg.in/olebedev/go-duktape.v3"
)

var bundle string

func init() {
	b, err := Asset("dist/bundle.js")
	must(err)
	bundle = string(b)
}

func Define(c *duktape.Context) {
	c.PushTimers()

	must(c.PevalString(`var process = process || {version: "v0.1"};`))
	must(c.PevalString(bundle))
	c.Pop2()
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
