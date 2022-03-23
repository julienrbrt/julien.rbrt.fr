package page

import (
	"time"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/julienrbrt/julien.rbrt.fr/component"
)

func (p *CV) renderProjects() *vecty.HTML {
	return elem.Span(
		vecty.Markup(vecty.Class("cv-entry")),
		elem.UnorderedList(
			&component.Experience{
				BeginDate:   time.Date(2021, time.January, 15, 0, 0, 0, 0, time.UTC),
				Location:    "Netherlands",
				JobTitle:    "Founder",
				Company:     "WoningFinder",
				Description: "Automatically react to houses from housing corporation in the Netherlands. Built with Go and Nuxt.js.",
			},
		),
	)
}
