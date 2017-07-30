# go-duktape-cheerio

Load [Cheerio](https://github.com/cheeriojs/cheerio) into [go-duktape](https://github.com/olebedev/go-duktape).

### Usage

First of all install the package `go get gopkg.in/olebedev/go-duktape-fetch.v3`.

```go
package main

import (
  "fmt"

  "gopkg.in/olebedev/go-duktape.v3"
  "github.com/cardigann/go-duktape-cheerio"
)

func main() {
  // create an ecmascript context
  ctx := duktape.New()

  // push cheerio into the global scope
  cheerio.Define(ctx)

  // use cheerio
  ctx.PevalString(`
    const cheerio = require('cheerio')
    const $ = cheerio.load('<h2 class="title">Hello world</h2>')

    $('h2.title').text('Hello there!')
    $('h2').addClass('welcome')

    $.html()
  `)

  // outputs "<h2 class="title welcome">Hello there!</h2>"
  fmt.Println(ctx.SafeToString(-1))
}
```

