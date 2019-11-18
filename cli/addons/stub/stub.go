// Package stub implements the stub addon, a general-purpose addon that shows a
// modeline and supports pluggable binding.
package stub

import (
	"github.com/elves/elvish/cli"
	"github.com/elves/elvish/cli/el"
	"github.com/elves/elvish/cli/el/layout"
	"github.com/elves/elvish/cli/term"
	"github.com/elves/elvish/edit/ui"
)

// Config keeps the configuration for the stub addon.
type Config struct {
	// Keybinding.
	Binding el.Handler
	// Name to show in the modeline.
	Name string
	// Whether the addon widget gets the focus.
	Focus bool
}

type widget struct {
	Config
}

func (w *widget) Render(width, height int) *ui.Buffer {
	buf := ui.NewBufferBuilder(width).
		WriteStyled(layout.ModeLine(w.Name, false)).SetDotHere().Buffer()
	buf.TrimToLines(0, height)
	return buf
}

func (w *widget) Handle(event term.Event) bool {
	return w.Binding.Handle(event)
}

func (w *widget) Focus() bool {
	return w.Config.Focus
}

// Start starts the stub addon.
func Start(app cli.App, cfg Config) {
	if cfg.Binding == nil {
		cfg.Binding = el.DummyHandler{}
	}
	w := widget{cfg}
	app.MutateState(func(s *cli.State) { s.Addon = &w })
	app.Redraw()
}
