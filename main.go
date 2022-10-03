package main

import (
	"bytes"
	"syscall/js"

	"github.com/leonhfr/polyglot/data"
	"github.com/leonhfr/polyglot/pkg/polyglot"
	"github.com/leonhfr/polyglot/pkg/position"
)

var book = polyglot.New()

func main() {
	js.Global().Set("polyglotBounds", bounds())
	js.Global().Set("polyglotLookup", lookup())
	<-make(chan struct{})
}

func init() {
	_ = book.Init(bytes.NewReader(data.Performance))
}

func bounds() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		min, max := book.Bounds()
		return []any{min, max}
	})
}

func lookup() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return "invalid number of arguments passed"
		}
		fen := args[0].String()
		pos := position.FromFEN(fen)
		res := make(map[string]any)
		for _, wm := range book.Lookup(pos) {
			res[wm.Move.String()] = wm.Weight
		}
		return res
	})
}
