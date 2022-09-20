package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/knipferrc/bubbletea-starter/internal/util/qr"
)

// View returns a string representation of the entire application UI.
func (b Bubble) View() string {
	var currentView string

	if b.showQR {
		return qr.Generate(b.textQR)
	}

	if !b.ready {
		return fmt.Sprintf("%s%s", b.loader.View(), "loading...")
	}

	if b.help.ShowAll {
		currentView = b.help.View(b.keys)
	} else {
		s := fmt.Sprintf("Welcome to the bubbletea-starter app")
		b.viewport.SetContent(s)

		currentView = b.viewport.View()
	}

	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Bold(true).
		Italic(true).
		Render(currentView)
}
