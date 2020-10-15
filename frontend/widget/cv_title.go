package widget

import (
	"fmt"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

// TitleWidget represents the widget that display the CV Title
type TitleWidget struct {
	vecty.Core
	Title string
}

func (w *TitleWidget) Render() vecty.ComponentOrHTML {
	return elem.Heading2(
		vecty.Markup(vecty.Class("cv-title")),
		vecty.Text(fmt.Sprintf("— %s", w.Title)),
	)
}
