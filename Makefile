build:
	GOOS=js GOARCH=wasm go build -o public/main.wasm ./cmd/render

.PHONY: build