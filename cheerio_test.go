package cheerio

import (
	"fmt"
	"testing"

	duktape "gopkg.in/olebedev/go-duktape.v3"
)

func TestAddTextAndClass(t *testing.T) {
	ctx := duktape.New()
	Define(ctx)

	defer ctx.DestroyHeap()

	js := `
		$ = cheerio.load('<h2 class = "title">Hello world</h2>');

		$('h2.title').text('Hello there!');
		$('h2').addClass('welcome');

		$.html();
	`

	if err := ctx.PevalString(js); err != nil {
		t.Fatal(err)
	}

	respString := ctx.SafeToString(-1)
	expect := `<html><head></head><body><h2 class="title welcome">Hello there!</h2></body></html>`

	if respString != expect {
		t.Fatalf("Expected %s, got %s", expect, respString)
	}
}

func TestGlobals(t *testing.T) {
	testCases := []struct {
		js     string
		typeOf string
	}{
		{`typeof cheerio;`, "function"},
	}
	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("#%d", idx), func(t *testing.T) {
			ctx := duktape.New()
			Define(ctx)
			defer ctx.Destroy()

			if err := ctx.PevalString(tc.js); err != nil {
				t.Fatal(err)
			}

			got := ctx.SafeToString(-1)
			if got != tc.typeOf {
				t.Fatalf("Expected typeOf of %s, got %s", tc.typeOf, got)
			}
		})
	}
}
