# julien.rbrt.fr - [live](https://julien.rbrt.fr)

## Build

```sh
cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" public/
GOOS=js GOARCH=wasm go build -buildvcs=false -o public/main.wasm ./cmd/render
```

## Notes

- Built with [Vecty](https://github.com/hexops/vecty)
- Using [Go WebAssembly](https://github.com/golang/go/wiki/WebAssembly)
- Architecture inspired from [marwan.io](https://github.com/marwan-at-work/marwanio)
