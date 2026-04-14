
.PHONY: all
all:
	exit 0

.PHONY: gojs
gojs:
	mkdir -p src/dist
	GOOS=js GOARCH=wasm go build -tags js -o src/dist/gojs.wasm ./cmd/gojs

# Фронт в docs/ для GitHub Pages (нужны yarn, mage)
.PHONY: site
site:
	mage Site

.PHONY: start
start: gojs
	yarn start

.PHONY: generate
generate:
	go install golang.org/x/tools/cmd/stringer@latest
	go generate ./...