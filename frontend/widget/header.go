package widget

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
)

// HeaderWidget is the main text in the top of the page.
type HeaderWidget struct {
	vecty.Core
	Title string
}

// Render renders the name in the top of the screen.
func (w *HeaderWidget) Render() vecty.ComponentOrHTML {
	return elem.Header(
		elem.Heading1(
			vecty.Markup(vecty.Class("logo")),
			elem.Anchor(
				vecty.Markup(prop.Href("//julien.rbrt.fr")),
				vecty.Text(w.Title)),
		),
	)
}
