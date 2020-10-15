package page

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"julien.rbrt.fr/frontend/widget"
)

// CV is the CV page.
type CV struct {
	vecty.Core
}

// Render implements the vecty.Component interface.
func (p *CV) Render() vecty.ComponentOrHTML {
	return elem.Div(
		&widget.HeaderWidget{Title: "Julien Robert"},
		elem.Section(
			vecty.Markup(vecty.Class("main")),
			// Experience Section
			&widget.TitleWidget{Title: "Experience"},
			p.renderExperience(),
			// Skills Section
			&widget.TitleWidget{Title: "Skills"},
			p.renderSkills(),
			// Education Section
			&widget.TitleWidget{Title: "Education"},
			p.renderEducation(),
			// Certifications Section
			&widget.TitleWidget{Title: "Certifications"},
			p.renderCertification(),
			// Interest Section
			&widget.TitleWidget{Title: "Interest"},
			p.renderInterest(),
		),
		&widget.FooterWidget{},
	)
}
