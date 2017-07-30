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

	js := `
		var process = process || {version: "v0.1"};

		Duktape.errCreate = function (e) {
			print(e.stack);
			return e;
		}
	`

	must(c.PevalString(js))
	must(c.PevalString(bundle))

	c.Pop()
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
