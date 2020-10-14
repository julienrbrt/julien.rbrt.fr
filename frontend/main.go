package main

import (
	"github.com/hexops/vecty"
	"julien.rbrt.fr/frontend/page"
)

func main() {
	vecty.SetTitle("julien.rbrt.fr")
	vecty.RenderBody(&page.About{})
}