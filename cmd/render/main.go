package main

import (
	"github.com/hexops/vecty"
	"github.com/julienrbrt/julien.rbrt.fr/page"
)

func main() {
	vecty.SetTitle("julien.rbrt.fr")
	vecty.RenderBody(&page.Body{})
}
