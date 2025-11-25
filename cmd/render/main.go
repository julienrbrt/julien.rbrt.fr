package main

import (
	"github.com/hexops/vecty"

	"go.rbrt.fr/julien.rbrt.fr/page"
)

func main() {
	vecty.SetTitle("julien.rbrt.fr")
	vecty.RenderBody(&page.Body{})
}
