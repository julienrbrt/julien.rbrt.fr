package main

import (
	"github.com/hexops/vecty"

	"pkg.rbrt.fr/julien.rbrt.fr/page"
)

func main() {
	vecty.SetTitle("julien.rbrt.fr")
	vecty.RenderBody(&page.Body{})
}
