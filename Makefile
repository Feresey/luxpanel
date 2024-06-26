
.PHONY: all
all:
	exit 0

.PHONY: gojs
gojs:
	GOOS=js GOARCH=wasm go build -o site/code/gojs.wasm ./cmd/gojs

generate:
	go install golang.org/x/tools/cmd/stringer@latest
	go generate ./...