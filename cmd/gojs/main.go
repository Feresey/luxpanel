//go:build js

package main

import "syscall/js"

func parseCombatLog(this js.Value, args []js.Value) any {
	return "hello world"
}

func parseGameLog(this js.Value, args []js.Value) any {
	return "hello world"
}

func main() {
	js.Global().Set("exportFunc", js.FuncOf(exportFunc))
	select {}
}
