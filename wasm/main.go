package main

import (
	"syscall/js"

	"github.com/caiguanhao/opencc"
)

func convert(this js.Value, i []js.Value) interface{} {
	if len(i) != 2 {
		return js.Undefined()
	}
	return js.ValueOf(opencc.Convert(i[0].String(), i[1].String()))
}

func main() {
	c := make(chan struct{}, 0)
	dicts := map[string]interface{}{}
	for name, text := range opencc.Dictionaries {
		dicts[name] = text
	}
	js.Global().Set("OpenCC", map[string]interface{}{
		"Dictionaries": dicts,
		"Convert":      js.FuncOf(convert),
	})
	js.Global().Call("eval", "document.dispatchEvent(new Event('OpenCCInited'))")
	<-c
}
