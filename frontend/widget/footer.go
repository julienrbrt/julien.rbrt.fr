package widget

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
)

// FooterWidget is the footer of the website
type FooterWidget struct {
	vecty.Core
}

// Render renders the footer
func (w *FooterWidget) Render() vecty.ComponentOrHTML {
	return elem.Footer(
		vecty.Markup(vecty.Class("footer")),
		elem.Div(
			vecty.Text("This website is written in Vecty - "),
			elem.Anchor(
				vecty.Markup(prop.Href("https://www.github.com/julienrbrt/julien.rbrt.fr")),
				vecty.Text("(source)"),
			),
		),
		elem.Span(
			vecty.Text("Hosted on DigitalOcean"),
		),
	)
}
