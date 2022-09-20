package tui

import (
	"github.com/knipferrc/bubbletea-starter/internal/config"
	"github.com/knipferrc/bubbletea-starter/internal/repository"
	"github.com/knipferrc/bubbletea-starter/internal/wa"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
)

// Bubble represents the state of the UI.
type Bubble struct {
	keys         keyMap
	help         help.Model
	loader       spinner.Model
	chatViewport viewport.Model
	appConfig    config.Config
	ready        bool

	chatState    stateView
	chatList     list.Model
	chatTextarea textarea.Model

	showQR bool
	textQR string
	client *wa.Client
}

// NewBubble creates an instance of the UI.
func NewBubble(cfg config.Config, client *wa.Client, db *repository.DB) Bubble {
	keys := getDefaultKeyMap()

	l := spinner.New()
	l.Spinner = spinner.Dot

	h := help.New()
	h.Styles.FullKey.Foreground(lipgloss.Color("#ffffff"))
	h.Styles.FullDesc.Foreground(lipgloss.Color("#ffffff"))

	return Bubble{
		keys:      keys,
		help:      h,
		loader:    l,
		appConfig: cfg,
		ready:     false,

		chatState:    listView,
		chatTextarea: defaultTextarea(),
		chatViewport: viewport.Model{},
		chatList:     list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0),

		client: client,
	}
}

type stateView uint

const (
	listView stateView = iota
	viewportView
	textareaView
)

func defaultTextarea() textarea.Model {
	ta := textarea.New()
	ta.Placeholder = "Send a message..."
	ta.Focus()

	ta.Prompt = "┃ "
	ta.CharLimit = 280

	// Remove cursor line styling
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()

	ta.ShowLineNumbers = false
	ta.KeyMap.InsertNewline.SetEnabled(false)

	return ta
}