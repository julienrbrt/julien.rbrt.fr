package page

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"julien.rbrt.fr/frontend/widget"
)

func (p *CV) renderSkills() *vecty.HTML {
	return elem.Span(
		vecty.Markup(vecty.Class("cv-entry")),
		elem.UnorderedList(
			&widget.ListItemWidget{
				Prefix: "General:",
				Title:  "Computer Science, Mobile Development, Backend Development, Data Analysis",
			},
			&widget.ListItemWidget{
				Prefix: "Programming:",
				Title:  "Dart (Flutter), Go, R, SQL",
			},
			&widget.ListItemWidget{
				Prefix: "Tools & Techniques:",
				Title:  "REST, Docker, PostgreSQL, MariaDB, Redis",
			},
			&widget.ListItemWidget{
				Prefix: "OS:",
				Title:  "GNU/Linux (Debian, Fedora, Ubuntu), Apple macOS",
			},
			&widget.ListItemWidget{
				Prefix: "Business:",
				Title:  "Agile, SCRUM, Kanban, Prince 2, UML, BPMN, Archimate",
			},
			&widget.ListItemWidget{
				Prefix: "Languages:",
				Title:  "French (Native), English (Fluent), Dutch (Working proficiency)",
			},
		),
	)
}
