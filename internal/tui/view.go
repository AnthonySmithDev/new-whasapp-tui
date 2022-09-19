package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/knipferrc/bubbletea-starter/internal/util/qr"
)

// View returns a string representation of the entire application UI.
func (b Bubble) View() string {
	var currentView string

	qr.Generate("https://www.google.com/search?q=bubbletea+center+text+github&sxsrf=ALiCzsYLXbs4_EQ0HSNwlScA8gmqTf8lXA%3A1663612342571&ei=trUoY5e9IpWL0AbW176ABQ")

	if !b.ready {
		return fmt.Sprintf("%s%s", b.loader.View(), "loading...")
	}

	if b.help.ShowAll {
		currentView = b.help.View(b.keys)
	} else {
		s := fmt.Sprintf("Welcome to the bubbletea-starter app\n Events received: %d", b.responses)
		b.viewport.SetContent(s)

		currentView = b.viewport.View()
	}

	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Bold(true).
		Italic(true).
		Render(currentView)
}
