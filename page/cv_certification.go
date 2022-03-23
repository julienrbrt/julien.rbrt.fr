package page

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/julienrbrt/julien.rbrt.fr/component"
)

func (p *CV) renderCertification() *vecty.HTML {
	return elem.Span(
		vecty.Markup(vecty.Class("cv-entry")),
		elem.UnorderedList(
			&component.ListItemWithURL{
				Title: "Data Science Professional Certificate",
				Affix: "HarvardX, edX",
			},
			&component.ListItemWithURL{
				Title: "R Programmer",
				Affix: "DataCamp",
			},
			&component.ListItemWithURL{
				Prefix: "15.071x:",
				Title:  "The Analytics Edge",
				Affix:  "MITx, edX",
				URL:    "https://courses.edx.org/certificates/eeeccf40dd1c43da8550b3705c26f1a2",
			},
			&component.ListItemWithURL{
				Prefix: "DS101x:",
				Title:  "Data Science Ethics",
				Affix:  "MichiganX, edX",
				URL:    "https://courses.edx.org/certificates/d70921526695401abf28f936f3e3fa9d",
			},
		),
	)
}
