package main

import (
	"bytes"
	"fmt"
	"strings"
	"syscall/js"

	annotations "github.com/rikatz/ingress-nginx-annotations"
	"github.com/sahilm/fuzzy"
)

var data = annotations.NewAnnotationFactory()

var keys []string

func main() {
	// Collect keys
	keys = make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}

	doc := js.Global().Get("document")
	input := doc.Call("getElementById", "q")
	out := doc.Call("getElementById", "out")

	cb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		q := strings.TrimSpace(input.Get("value").String())
		if q == "" {
			out.Set("innerHTML", "<i>Type something...</i>")
			return nil
		}

		matches := fuzzy.Find(q, keys)
		if len(matches) == 0 {
			out.Set("innerHTML", `<b>No match</b>`)
			return nil
		}

		var buf bytes.Buffer

		/*
					 <div class="w3-panel w3-white w3-card w3-display-container">
			   <p class="w3-text-blue"><b>email.zip</b></p>
			   <p>https://www.w3schools.com/lib/email.zip</p>
			 </div>
		*/
		for _, v := range matches {
			buf.WriteString(`<div class="w3-panel w3-white w3-card w3-display-container">`)
			val := data[v.Str]
			buf.WriteString(fmt.Sprintf("<p class=\"w3-text-blue\">%s</p>", v.Str))
			buf.WriteString(fmt.Sprintf("<p>%s</p>", val.GatewayAPI))
			buf.WriteString(`</div>`)

		}

		out.Set("innerHTML", buf.String())
		return nil
	})

	input.Call("addEventListener", "input", cb)
	out.Set("innerHTML", "<i>Type something...</i>")

	select {}
}

func htmlEscape(s string) string {
	replacer := strings.NewReplacer(
		"&", "&amp;",
		"<", "&lt;",
		">", "&gt;",
		`"`, "&quot;",
		"'", "&#39;",
	)
	return replacer.Replace(s)
}
