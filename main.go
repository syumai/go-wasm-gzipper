package main

import (
	"bytes"
	"io"
	"syscall/js"

	"github.com/syumai/go-wasm-gzipper/compressor"
)

func main() {
	js.Global().Set("compress", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) < 1 {
			panic("src must be given")
		}
		jsSrc := args[0]
		srcLen := jsSrc.Get("length").Int()
		srcBytes := make([]byte, srcLen)
		js.CopyBytesToGo(srcBytes, jsSrc)

		src := bytes.NewReader(srcBytes)

		r, err := compressor.Compress(src)
		if err != nil {
			panic(err)
		}

		var buf bytes.Buffer
		if _, err := io.Copy(&buf, r); err != nil {
			panic(err)
		}
		ua := newUint8Array(buf.Len())
		js.CopyBytesToJS(ua, buf.Bytes())
		return ua
	}))
	select {}
}
