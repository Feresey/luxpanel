//go:build js

package main

import (
	"context"
)

func main() {
	ctx := context.Background()

	r := NewRuntime(ctx)

	if err := r.Start(ctx); err != nil {
		panic(err)

	}
	if err := r.InitI18n(ctx); err != nil {
		panic(err)
	}
	r.RegisterJSBindings(ctx)

	select {}
}
