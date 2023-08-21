package main

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type MarketConfirm struct {
	help     help.Model
	keys     KeyMap
	input    textinput.Model
	quitting bool
}

func NewMarketConfirm() MarketConfirm {
	input := NewTextInput("please enter the exchange market")
	input.SetValue(opts.token0.Symbol() + "/" + opts.token1.Symbol())
	return MarketConfirm{
		help:  help.New(),
		input: input,
		keys: KeyMap{
			keys: map[KeyType]key.Binding{
				KeyEnter: NewEnterBinding("confirm"),
				KeyQuit:  KeyBindingQuit,
			},
			short: KeyEnter | KeyQuit,
		},
	}
}

func (m MarketConfirm) Init() tea.Cmd {
	return nil
}

func (m MarketConfirm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.keys[KeyEnter]):
			opts.market = m.input.Value()
			m.quitting = true

			if opts.market == "" {
				return m, tea.Quit
			}
			return NewExchangeSelector(), tea.Println(
				green.Render("Market: "),
				grey.Render(opts.market),
			)
		case key.Matches(msg, m.keys.keys[KeyQuit]):
			m.quitting = true
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m MarketConfirm) View() string {
	if m.quitting {
		if opts.market == "" {
			return red.Render("Exchange market can not be empty!!!\n")
		}
		return ""
	}
	s := blue.Render("Please confirm whether the market is correct?") + "\n"
	s += m.input.View()
	s += padding.Render("\n\n" + m.help.View(m.keys) + "\n")
	return s
}
