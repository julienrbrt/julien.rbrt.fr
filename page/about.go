package page

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
	"github.com/julienrbrt/julien.rbrt.fr/component"
)

// About is the main page.
type About struct {
	vecty.Core
}

// Render implements the vecty.Component interface.
func (p *About) Render() vecty.ComponentOrHTML {
	return elem.Div(
		&component.Header{Title: "julien.rbrt.fr"},
		elem.Section(
			vecty.Markup(vecty.Class("main")),
			elem.Heading1(vecty.Text("I am "),
				elem.Anchor(
					vecty.Markup(prop.Href("cv")),
					vecty.Text("Julien Robert"),
				),
				vecty.Text(", Blockchain Engineer, currently based in the Netherland.")),
			elem.Heading1(vecty.Text("Go and Crypto enthusiast, explore my contributions on "),
				elem.Anchor(
					vecty.Markup(prop.Href("https://github.com/julienrbrt")),
					vecty.Text("GitHub"),
				),
				vecty.Text(" or find me on "),
				elem.Anchor(
					vecty.Markup(prop.Href("https://x.com/_julienrbrt")),
					vecty.Text("𝕏"),
				), vecty.Text("."),
			),
			p.renderContact(),
		),
		&component.Footer{},
	)
}

func (p *About) renderContact() *vecty.HTML {
	return elem.Heading1(vecty.Text("To contact me directly, send an email to:"),
		elem.Preformatted(
			vecty.Markup(vecty.Class("email")),
			vecty.Text("julien at rbrt dot fr"),
		),
	)
}
