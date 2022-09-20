package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/knipferrc/bubbletea-starter/internal/util/qr"
)

var (
	defaultStyle = lipgloss.NewStyle().
		// Padding(1, 2).
		BorderStyle(lipgloss.NormalBorder())
	focusedStyle = defaultStyle.Copy().
			BorderForeground(lipgloss.Color("69"))
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
		listRender := defaultStyle.Render(b.chatList.View())
		viewportRender := defaultStyle.Render(b.chatViewport.View())
		textareaRender := defaultStyle.Render(b.chatTextarea.View())

		switch b.chatState {
		case listView:
			listRender = focusedStyle.Render(b.chatList.View())
		case viewportView:
			viewportRender = focusedStyle.Render(b.chatViewport.View())
		case textareaView:
			textareaRender = focusedStyle.Render(b.chatTextarea.View())
		}

		currentView = lipgloss.JoinHorizontal(lipgloss.Left, listRender,
			lipgloss.JoinVertical(
				lipgloss.Top,
				viewportRender,
				textareaRender,
			),
		)
	}
	return currentView
}
