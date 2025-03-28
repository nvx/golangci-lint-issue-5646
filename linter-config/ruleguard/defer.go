//go:build ruleguard

package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func errorOnPrintln(m dsl.Matcher) {
	m.Import("fmt")
	m.Match(`fmt.Println($*args)`).
		Report(`contrived example to prevent fmt.Println'`)
}
