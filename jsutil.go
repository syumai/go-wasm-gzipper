package main

import "syscall/js"

var global = js.Global()

func newUint8Array(size int) js.Value {
	ua := global.Get("Uint8Array")
	return ua.New(size)
}
