package component

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
)

// Footer is the footer of the website
type Footer struct {
	vecty.Core
}

// Render renders the footer
func (w *Footer) Render() vecty.ComponentOrHTML {
	return elem.Footer(
		vecty.Markup(vecty.Class("footer")),
		elem.Div(
			vecty.Text("This website is using Vecty - "),
			elem.Anchor(
				vecty.Markup(prop.Href("https://www.github.com/julienrbrt/julien.rbrt.fr")),
				vecty.Text("(source)"),
			),
		),
	)
}
