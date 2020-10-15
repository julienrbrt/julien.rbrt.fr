package page

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"julien.rbrt.fr/frontend/widget"
)

func (p *CV) renderCertification() *vecty.HTML {
	return elem.Span(
		vecty.Markup(vecty.Class("cv-entry")),
		elem.UnorderedList(
			&widget.ListItemWithURLWidget{
				Title: "Data Science Professional Certificate",
				Affix: "HarvardX, edX",
			},
			&widget.ListItemWithURLWidget{
				Title: "R Programmer",
				Affix: "DataCamp",
			},
			&widget.ListItemWithURLWidget{
				Prefix: "15.071x:",
				Title:  "The Analytics Edge",
				Affix:  "MITx, edX",
				URL:    "https://courses.edx.org/certificates/eeeccf40dd1c43da8550b3705c26f1a2",
			},
			&widget.ListItemWithURLWidget{
				Prefix: "DS101x:",
				Title:  "Data Science Ethics",
				Affix:  "MichiganX, edX",
				URL:    "https://courses.edx.org/certificates/d70921526695401abf28f936f3e3fa9d",
			},
		),
	)
}
