package page

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"

	"pkg.rbrt.fr/julien.rbrt.fr/component"
)

func (p *CV) renderTalks() *vecty.HTML {
	return elem.Span(
		vecty.Markup(vecty.Class("cv-entry")),
		elem.UnorderedList(
			&component.ListItemWithURL{
				Title: "Evolve your chain",
				Affix: "Cosmoverse 25",
				URL:   "https://www.youtube.com/watch?v=avMTVTJmtSQ",
			},
			&component.ListItemWithURL{
				Title: "Build Your Own Chain with Cosmos SDK",
				Affix: "BUIDL",
				URL:   "https://www.youtube.com/watch?v=Pr698TNMSXo",
			},
			&component.ListItemWithURL{
				Title: "The Cosmos SDK v2",
				Affix: "Cosmoverse 24",
				URL:   "https://www.youtube.com/watch?v=Tjg8EDw6sSQ",
			},
			&component.ListItemWithURL{
				Title: "Smart Accounts",
				Affix: "Cosmoverse 24",
				URL:   "https://www.youtube.com/watch?v=MJf_tWh0ZM0",
			},
			&component.ListItemWithURL{
				Title: "Beyond Scaffolding Ignite",
				Affix: "Nebular 24",
				URL:   "https://www.youtube.com/watch?v=aLywSzhZCac",
			},
			&component.ListItemWithURL{
				Title: "Introduction to app.go",
				Affix: "Interchain Academy 23",
				URL:   "https://www.youtube.com/watch?v=G6QUIUwYaSU",
			},
		),
	)
}
