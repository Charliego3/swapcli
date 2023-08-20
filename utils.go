package main

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func NewTextInput(placeholder string) textinput.Model {
	ti := textinput.New()
	ti.Placeholder = placeholder
	ti.Focus()
	ti.CharLimit = 200
	ti.Width = 60
	ti.PromptStyle = optionStyle.Copy().Bold(true)
	ti.TextStyle = green
	return ti
}

func NewHelpBinding() key.Binding {
	return key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	)
}

func NewEnterBinding(desc string) key.Binding {
	return key.NewBinding(
		key.WithKeys(tea.KeyEnter.String()),
		key.WithHelp("\U000F17A5", desc),
	)
}

func NewSelectKeyMap() KeyMap {
	return KeyMap{
		keys: map[KeyType]key.Binding{
			KeyUp: key.NewBinding(
				key.WithKeys("up", "k", "K"),
				key.WithHelp("↑/k", "move up"),
			),
			KeyDown: key.NewBinding(
				key.WithKeys("down", "j", "J"),
				key.WithHelp("↓/j", "move down"),
			),
			KeyEnter: NewEnterBinding("choose"),
			KeyHelp:  NewHelpBinding(),
			KeyQuit:  KeyBindingQuit,
		},
		short: KeyHelp | KeyQuit,
		full: []KeyType{
			KeyUp | KeyDown | KeyEnter,
			KeyHelp | KeyQuit,
		},
	}
}
