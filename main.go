package main

import (
	"bytes"
	"io"
	"syscall/js"

	"github.com/syumai/go-wasm-gzipper/compressor"
)

func newCompressor(this js.Value, args []js.Value) interface{} {
	name := args[0].String()
	size := args[1].Int()
	c := compressor.New(name, size)
	this.Set("buffer", js.TypedArrayOf(c.Buffer()))
	this.Set("compress", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		r, err := c.Compress()
		if err != nil {
			panic(err)
		}

		var buf bytes.Buffer
		if _, err := io.Copy(&buf, r); err != nil {
			panic(err)
		}
		return js.TypedArrayOf([]uint8(buf.Bytes()))
	}))
	return this
}

func main() {
	window := js.Global()
	window.Set("Compressor", js.FuncOf(newCompressor))
	select {}
}
