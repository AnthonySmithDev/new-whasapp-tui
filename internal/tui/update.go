package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/knipferrc/bubbletea-starter/internal/util/qr"
	"go.mau.fi/whatsmeow"
)

// Update handles updating the UI.
func (b Bubble) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case responseMsg:
		b.responses++                    // record external activity
		return b, waitForActivity(b.Sub) // wait for next event

	case whatsmeow.QRChannelItem:
		fmt.Print(qr.Generate(msg.Code))
		return b, waitForQR(b.client.QRChannel)

	case tea.WindowSizeMsg:
		// qrStyle = qrStyle.Width(msg.Width).Height(msg.Height)
		b.viewport.Height = msg.Height
		b.viewport.Width = msg.Width
		b.help.Width = msg.Width

		if !b.ready {
			b.ready = true
		}

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, b.keys.Quit):
			return b, tea.Quit
		case key.Matches(msg, b.keys.Help):
			b.help.ShowAll = !b.help.ShowAll

			return b, nil
		}
	}

	b.loader, cmd = b.loader.Update(msg)
	cmds = append(cmds, cmd)

	b.viewport, cmd = b.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return b, tea.Batch(cmds...)
}
