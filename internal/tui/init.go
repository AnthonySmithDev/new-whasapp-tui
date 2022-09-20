package tui

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"go.mau.fi/whatsmeow"
)

type qrMsg whatsmeow.QRChannelItem

func waitForQR(ch <-chan whatsmeow.QRChannelItem) tea.Cmd {
	return func() tea.Msg {
		return qrMsg(<-ch)
	}
}

// Init initializes the UI.
func (b Bubble) Init() tea.Cmd {
	var cmds []tea.Cmd

	cmds = append(cmds,
		spinner.Tick,
		waitForQR(b.client.QRChannel),
	)

	return tea.Batch(cmds...)
}
