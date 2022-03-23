package page

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

func (p *CV) renderInterest() *vecty.HTML {
	return elem.Span(
		vecty.Markup(vecty.Class("cv-entry")),
		elem.UnorderedList(
			elem.ListItem(vecty.Text("Zero to One, by Peter Thiel")),
			elem.ListItem(vecty.Text("The Design of Everyday Things, by Don Norman")),
			elem.ListItem(vecty.Text("H2G2 series, by Douglas Adams")),
			elem.ListItem(vecty.Text("What If, by Randall Munroe")),
		),
	)
}
