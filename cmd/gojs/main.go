package main

import "syscall/js"

func exportFunc(this js.Value, args []js.Value) any {
	return "hello world"
}

func main() {
	js.Global().Set("exportFunc", js.FuncOf(exportFunc))
	select {}
}
