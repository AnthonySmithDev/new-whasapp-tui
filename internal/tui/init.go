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

type connectedMsg struct{}

func waitForConnected(ch <-chan struct{}) tea.Cmd {
	return func() tea.Msg {
		return connectedMsg(<-ch)
	}
}

// Init initializes the UI.
func (b Bubble) Init() tea.Cmd {
	var cmds []tea.Cmd

	cmds = append(cmds,
		spinner.Tick,
		waitForConnected(b.client.Connected),
		waitForQR(b.client.QRChannel),
	)

	return tea.Batch(cmds...)
}
