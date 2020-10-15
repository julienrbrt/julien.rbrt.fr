# julien.rbrt.fr - [live](https://julien.rbrt.fr)

## Development

To test the website locally run ``cd frontend; GOOS=js GOARCH=wasm go build -o ../public/main.wasm; cd ..; go run .``. This combine the two steps below.

### WASM
Build the `main.wasm`:
```
GOOS=js GOARCH=wasm go build -o ../public/main.wasm
```

### Server
Run the server with `go run .`

## Notes

* Built with [Vecty](https://github.com/hexops/vecty)
* Hosted on [DigitalOcean](https://www.digitalocean.com/products/app-platform/)
* Architecture inspired from [marwan.io](https://github.com/marwan-at-work/marwanio)