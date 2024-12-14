//go:build js

package main

import (
	"context"
	"errors"
	"fmt"
	"syscall/js"

	"net/http"
	_ "net/http/pprof"
)

func main() {
	ctx := context.Background()

	r := NewRuntime(ctx)

	if err := r.Start(ctx); err != nil {
		panic(err)

	}
	println("Starting")
	js.Global().Set("parseFiles", js.FuncOf(func(this js.Value, args []js.Value) any {
		defer println("exited")
		println("processing")
		if len(args) != 2 {
			println("exit on error")
			return fmt.Sprintf("2 arguments required, %d got", len(args))
		}
		if args[0].Type() != js.TypeString || args[1].Type() != js.TypeString {
			println("exit on error 2")
			return fmt.Sprintf("string arguments expected, but got: %s %s", args[0].Type(), args[1].Type())
		}
		if err := r.LoadFiles(ctx, args[0].String(), args[1].String()); err != nil {
			println("exit on error 3", err.Error())
			return fmt.Sprintf("LoadFiles: %v", err)
		}
		println("processed")
		return nil
	}))

	js.Global().Set("showLevels", js.FuncOf(func(this js.Value, args []js.Value) any {
		return len(r.Data.Levels)
	}))

	if err := http.ListenAndServe(":8888", http.DefaultServeMux); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			println(err.Error())
		}
		println("server stopped")
	}
	select {}
}
