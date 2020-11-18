package page

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
	"julien.rbrt.fr/frontend/widget"
)

// About is the main page.
type About struct {
	vecty.Core
}

// Render implements the vecty.Component interface.
func (p *About) Render() vecty.ComponentOrHTML {
	return elem.Div(
		&widget.HeaderWidget{Title: "julien.rbrt.fr"},
		elem.Section(
			vecty.Markup(vecty.Class("main")),
			elem.Heading1(vecty.Text("I am Julien Robert, Go backend software developer, currently living in the Netherlands")),
			elem.Heading1(vecty.Text("Dart, Go and data analysis enthusiast, you can find my code on "),
				elem.Anchor(
					vecty.Markup(prop.Href("https://github.com/julienrbrt")),
					vecty.Text("GitHub"),
				),
				vecty.Text(" or discover my side projects "),
				elem.Anchor(
					vecty.Markup(prop.Href("https://delice.app")),
					vecty.Text("Délice"),
				),
				vecty.Text(" and "),
				elem.Anchor(
					vecty.Markup(prop.Href("https://woningfinder.nl")),
					vecty.Text("WoningFinder"),
				),
				vecty.Text("."),
			),
			p.renderContact(),
		),
		&widget.FooterWidget{},
	)
}

func (p *About) renderContact() *vecty.HTML {
	return elem.Heading1(vecty.Text("For contacting me directly send an email at:"),
		elem.Preformatted(
			vecty.Markup(vecty.Class("email")),
			vecty.Text("julien at rbrt dot fr"),
		),
	)
}
