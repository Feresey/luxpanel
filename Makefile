
.PHONY: all
all:
	exit 0

.PHONY: gojs
gojs:
	GOOS=js GOARCH=wasm go build -o mysite/code/gojs.wasm ./cmd/gojs
	mkdir -p src/dist && cp mysite/code/gojs.wasm src/dist/gojs.wasm

.PHONY: start
start: gojs
	yarn start

.PHONY: generate
generate:
	go install golang.org/x/tools/cmd/stringer@latest
	go generate ./...