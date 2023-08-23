package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"swapcli/apis"
)

type ExchangeSelector struct {
	keyMap KeyMap
	help   help.Model
	key0   textinput.Model
	key1   textinput.Model
	cursor int
	done   bool
}

func NewExchangeSelector() *ExchangeSelector {
	selector := &ExchangeSelector{
		keyMap: NewSelectKeyMap(),
		help:   help.New(),
		key0:   NewTextInput("please enter access key"),
		key1:   NewTextInput("please enter secret key"),
	}
	selector.key1.Blur()
	return selector
}

func (e *ExchangeSelector) Init() tea.Cmd {
	return nil
}

func (e *ExchangeSelector) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch k := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(k, e.keyMap.keys[KeyEnter]):
			if opts.exchange == nil {
				opts.exchange = apis.Fetch(apis.ExchangeTypeValues()[e.cursor])
				e.keyMap.keys[KeyEnter] = NewEnterBinding("confirm")
				e.keyMap.keys[KeyTab] = key.NewBinding(
					key.WithKeys(tea.KeyTab.String()),
					key.WithHelp("\U000F0312", "toggle focus"),
				)
				e.keyMap.keys[KeyUp] = key.NewBinding(
					key.WithKeys(tea.KeyUp.String()),
					key.WithHelp("↑", "focus up"),
				)
				e.keyMap.keys[KeyDown] = key.NewBinding(
					key.WithKeys(tea.KeyDown.String()),
					key.WithHelp("↓", "focus down"),
				)
				e.keyMap.full[2] = KeyTab | e.keyMap.full[2]
				return e, tea.Batch(tea.Println(green.Render("Exchange: "),
					grey.Render(apis.ExchangeTypeValues()[e.cursor].String())), textinput.Blink)
			}

			accessFilled := e.key0.Value() != ""
			secretFilled := e.key1.Value() != ""
			if accessFilled && secretFilled {
				opts.accessKey = e.key0.Value()
				opts.secretKey = e.key1.Value()
				return e, tea.Quit
			}

			if accessFilled && !secretFilled && !e.key1.Focused() {
				e.key0.Blur()
				return e, e.key1.Focus()
			}
			if !accessFilled && secretFilled && !e.key0.Focused() {
				e.key1.Blur()
				return e, e.key0.Focus()
			}
		case key.Matches(k, e.keyMap.keys[KeyDown]):
			if cmd := e.toggleFocus(); cmd != nil {
				return e, cmd
			}

			e.cursor++
			if e.cursor > len(apis.ExchangeTypeValues())-1 {
				e.cursor = 0
			}
		case key.Matches(k, e.keyMap.keys[KeyUp]):
			if cmd := e.toggleFocus(); cmd != nil {
				return e, cmd
			}

			e.cursor--
			if e.cursor < 0 {
				e.cursor = len(apis.ExchangeTypeValues()) - 1
			}
		case key.Matches(k, e.keyMap.keys[KeyTab]):
			return e, e.toggleFocus()
		case key.Matches(k, e.keyMap.keys[KeyHelp]):
			e.help.ShowAll = !e.help.ShowAll
			return e, nil
		case key.Matches(k, e.keyMap.keys[KeyQuit]):
			e.done = true
			return e, tea.Quit
		}
	}

	var cmd0, cmd1 tea.Cmd
	e.key0, cmd0 = e.key0.Update(msg)
	e.key1, cmd1 = e.key1.Update(msg)
	return e, tea.Batch(cmd0, cmd1)
}

func (e *ExchangeSelector) toggleFocus() tea.Cmd {
	if opts.exchange == nil {
		return nil
	}

	var cmd tea.Cmd
	if e.key0.Focused() {
		e.key0.Blur()
		cmd = e.key1.Focus()
	} else {
		e.key1.Blur()
		cmd = e.key0.Focus()
	}
	return cmd
}

func (e *ExchangeSelector) View() string {
	if e.done {
		return "\n"
	}

	var s string
	if opts.exchange != nil {
		s = blue.Render("What's exchange api account?") + "\n"
		s += e.key0.View() + "\n"
		s += e.key1.View() + "\n"
	} else {
		s = blue.Render("What's exchange do you want to trade?") + "\n"
		for i, choice := range apis.ExchangeTypeValues() {
			if e.cursor == i {
				s += green.Render(fmt.Sprintf("\U000F012D %s", choice)) + "\n"
			} else {
				s += optionStyle.Render(fmt.Sprintf("  %s", choice)) + "\n"
			}
		}
	}

	s += padding.Render("\n" + e.help.View(e.keyMap) + "\n")
	return s
}
