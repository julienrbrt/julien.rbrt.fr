package page

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"

	"pkg.rbrt.fr/julien.rbrt.fr/component"
)

// CV is the CV page.
type CV struct {
	vecty.Core
}

// Render implements the vecty.Component interface.
func (p *CV) Render() vecty.ComponentOrHTML {
	return elem.Div(
		&component.Header{Title: "Julien Robert"},
		elem.Section(
			vecty.Markup(vecty.Class("main")),
			// Experience section
			&component.Title{Title: "Experience"},
			p.renderExperience(),
			// Side Projects section
			&component.Title{Title: "Side Projects"},
			p.renderProjects(),
			// Skills section
			&component.Title{Title: "Skills"},
			p.renderSkills(),
			// Education section
			&component.Title{Title: "Education"},
			p.renderEducation(),
			// Talks section
			&component.Title{Title: "Talks"},
			p.renderTalks(),
			// Certifications section
			&component.Title{Title: "Certifications"},
			p.renderCertification(),
		),
		&component.Footer{},
	)
}
