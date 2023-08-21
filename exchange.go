package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"swapcli/apis"
)

type ExchangeSelector struct {
	keyMap KeyMap
	help   help.Model
	cursor int
	done   bool
}

func NewExchangeSelector() *ExchangeSelector {
	return &ExchangeSelector{
		keyMap: NewSelectKeyMap(),
		help:   help.New(),
	}
}

func (e *ExchangeSelector) Init() tea.Cmd {
	return nil
}

func (e *ExchangeSelector) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch k := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(k, e.keyMap.keys[KeyEnter]):
			opts.exchange = apis.Fetch(apis.ExchangeTypeValues()[e.cursor])
			return e, tea.Quit
		case key.Matches(k, e.keyMap.keys[KeyDown]):
			e.cursor++
			if e.cursor > len(apis.ExchangeTypeValues())-1 {
				e.cursor = 0
			}
		case key.Matches(k, e.keyMap.keys[KeyUp]):
			e.cursor--
			if e.cursor < 0 {
				e.cursor = len(apis.ExchangeTypeValues()) - 1
			}
		case key.Matches(k, e.keyMap.keys[KeyHelp]):
			e.help.ShowAll = !e.help.ShowAll
		case key.Matches(k, e.keyMap.keys[KeyQuit]):
			e.done = true
			return e, tea.Quit
		}
	}
	return e, nil
}

func (e *ExchangeSelector) View() string {
	if opts.exchange != nil || e.done {
		return fmt.Sprintf("%T\n", opts.exchange)
	}

	s := blue.Render("What's exchange do you want to trade?") + "\n"
	for i, choice := range apis.ExchangeTypeValues() {
		if e.cursor == i {
			s += green.Render(fmt.Sprintf("\U000F012D %s", choice)) + "\n"
		} else {
			s += optionStyle.Render(fmt.Sprintf("  %s", choice)) + "\n"
		}
	}

	s += padding.Render("\n" + e.help.View(e.keyMap) + "\n")
	return s
}
