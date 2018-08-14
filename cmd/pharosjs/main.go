package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	core := js.Global.Get("document").Call("getElementById", "jot-core")
	core.Set("innerHTML", chroma())
	core.Get("style").Set("width", "800px")
	core.Get("style").Set("border", "dotted 3px gray")
	core.Get("style").Set("fontSize", "16px")
	core.Get("style").Set("fontFamily", "'Play', sans-serif")
	core.Get("style").Set("marginLeft", "auto")
	core.Get("style").Set("marginRight", "auto")
}

func chroma() string {
	r := bytes.NewReader([]byte(goSnippet()))
	var wb []byte
	w := bytes.NewBuffer(wb)
	lexer := lexers.Get("go")
	style := styles.Get("bw")
	formatter := html.New(html.WithLineNumbers(), html.TabWidth(3), html.LineNumbersInTable())
	contents, err := ioutil.ReadAll(r)
	iterator, err := lexer.Tokenise(nil, string(contents))
	err = formatter.Format(w, style, iterator)
	if err != nil {
		fmt.Println(err)
	}
	return string(w.Bytes())
}
