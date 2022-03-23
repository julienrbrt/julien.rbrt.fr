# julien.rbrt.fr - [live](https://julien.rbrt.fr)

## Build

```sh
GOOS=js GOARCH=wasm go build -o public/main.wasm ./cmd/render
```

## Notes

* Built with [Vecty](https://github.com/hexops/vecty)
* Using [Go WebAssembly](https://github.com/golang/go/wiki/WebAssembly)
* Architecture inspired from [marwan.io](https://github.com/marwan-at-work/marwanio)
