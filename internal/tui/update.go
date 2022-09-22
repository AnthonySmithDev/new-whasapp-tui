package tui

import (
	// "github.com/charmbracelet/bubbles/key"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// Update handles updating the UI.
func (b Bubble) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case qrMsg:
		if msg.Event == "code" {
			b.showQR = true
			b.textQR = msg.Code
			return b, waitForQR(b.client.QRChannel)
		} else if msg.Event == "success" {
			b.showQR = false
			b.textQR = ""
			return b, nil
		}

	case connectedMsg:
		b.ready = true
		return b, waitForConnected(b.client.Connected)

	case tea.WindowSizeMsg:
		h, v := defaultStyle.GetFrameSize()

		leftWidth := msg.Width*40/100 - h
		rigthWidth := msg.Width*60/100 - h

		listHeight := msg.Height - v

		viewportHeight := msg.Height - 5 - v
		textareaHeight := 5 - v
		if msg.Height < 20 {
			viewportHeight = msg.Height - 3 - v
			textareaHeight = 3 - v
		} else if msg.Height < 30 {
			viewportHeight = msg.Height - 4 - v
			textareaHeight = 4 - v
		}

		// set list size
		b.chatList.SetHeight(listHeight)
		b.chatList.SetWidth(leftWidth)

		//chat set viewport size
		b.chatViewport.Height = viewportHeight
		b.chatViewport.Width = rigthWidth

		//chat set textarea size
		b.chatTextarea.SetWidth(rigthWidth)
		b.chatTextarea.SetHeight(textareaHeight)

		b.help.Width = msg.Width
		b.chatViewport.SetContent("Welcome to the bubbletea-starter app")

	case tea.KeyMsg:
		// switch {
		// case key.Matches(msg, b.keys.Quit):
		// 	return b, tea.Quit
		// case key.Matches(msg, b.keys.Help):
		// 	b.help.ShowAll = !b.help.ShowAll

		// 	return b, nil
		// }

		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return b, tea.Quit
		case tea.KeyTab:
			switch b.chatState {
			case listView:
				b.chatState = viewportView
			case viewportView:
				b.chatState = textareaView
			case textareaView:
				b.chatState = listView
			}
		case tea.KeyEnter:
			switch b.chatState {
			case listView:
				b.messages = []string{}
				listItem := b.chatList.SelectedItem().(item)
				b.chatJID = listItem.GetID()
				messages := b.db.Message.FindMany(b.chatJID)
				b.messages = messages.ToList()
				b.chatViewport.SetContent(strings.Join(b.messages, "\n"))
				b.chatViewport.GotoBottom()
				b.chatState = textareaView
				break
			case viewportView:
				break
			case textareaView:
			}
		}

		switch b.chatState {
		case listView:
			b.chatList, cmd = b.chatList.Update(msg)
			cmds = append(cmds, cmd)
		case viewportView:
			b.chatViewport, cmd = b.chatViewport.Update(msg)
			cmds = append(cmds, cmd)
		case textareaView:
			b.chatTextarea, cmd = b.chatTextarea.Update(msg)
			cmds = append(cmds, cmd)
		}
	}

	b.loader, cmd = b.loader.Update(msg)
	cmds = append(cmds, cmd)

	return b, tea.Batch(cmds...)
}
