package tui

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
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

type messageMsg *events.Message

func waitForMessage(ch <-chan *events.Message) tea.Cmd {
	return func() tea.Msg {
		return messageMsg(<-ch)
	}
}

// Init initializes the UI.
func (b Bubble) Init() tea.Cmd {
	var cmds []tea.Cmd

	cmds = append(cmds,
		spinner.Tick,
		waitForQR(b.client.QRChannel),
		waitForConnected(b.client.Connected),
		waitForMessage(b.client.Message),
	)

	return tea.Batch(cmds...)
}
