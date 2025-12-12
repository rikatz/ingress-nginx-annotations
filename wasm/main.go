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

		for _, v := range matches {
			buf.WriteString(`<div class="w3-card-4">`)
			val := data[v.Str]
			gw := val.GatewayAPI
			icon := "fa-check-circle"
			theme := "w3-light-blue"
			if gw == "" {
				icon = "fa-exclamation-circle"
				gw = "Not supported yet."
				theme = "w3-pale-red"
			}
			buf.WriteString(fmt.Sprintf("<header class=\"w3-container %s\"><h1>%s <i class=\"fa %s\"></i></h1></header><div class=\"w3-container\">", theme, v.Str, icon))
			buf.WriteString(fmt.Sprintf("<p>%s</p>", gw))
			if val.GatewayAPIRef != "" {
				buf.WriteString(fmt.Sprintf("<p><a href=\"%s\" target=\"_blank\" rel=\"noopener noreferrer\">Reference</a></p>", val.GatewayAPIRef))
			}

			buf.WriteString(`</div></div>`)

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
