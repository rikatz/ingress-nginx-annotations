.DEFAULT_GOAL := pages

.PHONY: pages
pages:
	cp -f $(shell go env GOROOT)/lib/wasm/wasm_exec.js site/
	GOOS=js GOARCH=wasm go build -o site/main.wasm ./wasm/ 