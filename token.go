package main

import tea "github.com/charmbracelet/bubbletea"

type TokenSelector struct{}

func (t TokenSelector) Init() tea.Cmd {
	return nil
}

func (t TokenSelector) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyEnter {
			return t, tea.Quit
		}
	}

	return t, nil
}

func (t TokenSelector) View() string {
	return "token selector"
}
