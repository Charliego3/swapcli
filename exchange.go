package main

import "github.com/charmbracelet/bubbles/help"

type ExchangeSelector struct {
	keys    KeyMap
	help    help.Model
	cursor  int
	choices []string
}
