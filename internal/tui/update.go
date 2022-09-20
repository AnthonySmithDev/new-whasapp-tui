package tui

import (
	// "github.com/charmbracelet/bubbles/key"
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
		} else if msg.Event == "success" {
			b.showQR = false
			b.textQR = ""
		}
		return b, waitForQR(b.client.QRChannel)

	case tea.WindowSizeMsg:
		h, v := defaultStyle.GetFrameSize()

		leftWidth := msg.Width*40/100 - h
		rigthWidth := msg.Width*60/100 - h

		listHeight := msg.Height - v
		viewportHeight := msg.Height*90/100 - v
		textareaHeight := msg.Height*10/100 - v
		if msg.Height < 30 {
			viewportHeight = msg.Height*80/100 - v
			textareaHeight = msg.Height*20/100 - v
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

		if !b.ready {
			b.ready = true
		}

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
