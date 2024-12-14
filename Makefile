
.PHONY: all
all:
	exit 0

.PHONY: gojs
gojs:
	GOOS=js GOARCH=wasm go build -o src/dist/gojs.wasm ./cmd/gojs

.PHONY: start
start: gojs
	yarn start

.PHONY: generate
generate:
	go install golang.org/x/tools/cmd/stringer@latest
	go generate ./...