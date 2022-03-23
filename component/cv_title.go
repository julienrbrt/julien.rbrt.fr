package component

import (
	"fmt"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

// Title represents the component that display the CV Title
type Title struct {
	vecty.Core
	Title string
}

func (w *Title) Render() vecty.ComponentOrHTML {
	return elem.Heading2(
		vecty.Markup(vecty.Class("cv-title")),
		vecty.Text(fmt.Sprintf("— %s", w.Title)),
	)
}
