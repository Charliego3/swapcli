package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/daoleno/uniswap-sdk-core/entities"
	"os"
	"strconv"
)

type Options struct {
	chainId int
	token0  *entities.Token
	token1  *entities.Token
}

var opts = &Options{}

type App struct {
	cursor  int
	choices []string
}

func (c App) Init() tea.Cmd {
	return nil
}

func (c App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case tea.KeyEnter.String():
			chain := c.choices[c.cursor]
			if chain == "BNB Smart Chain MainNet" {
				opts.chainId = 56
			} else if chain == "Ethereum MainNet" {
				opts.chainId = 1
			}
			return TokenSelector{}, tea.Println(c.choices[c.cursor] + ", " + strconv.Itoa(opts.chainId))
		case tea.KeyDown.String(), "j":
			c.cursor++
			if c.cursor > len(c.choices)-1 {
				c.cursor = 0
			}
		case tea.KeyUp.String(), "k":
			c.cursor--
			if c.cursor < 0 {
				c.cursor = len(c.choices) - 1
			}
		case tea.KeyCtrlC.String(), "q":
			return c, tea.Quit
		}
	}

	return c, nil
}

func (c App) View() string {
	s := "What's network do you want to choose?\n\n"
	for i, choice := range c.choices {
		cursor := " " // no cursor
		if c.cursor == i {
			cursor = ">" // cursor!
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\nPress q to quit.\n"
	return s
}

func main() {
	app := tea.NewProgram(App{
		choices: []string{
			"BNB Smart Chain MainNet",
			"Ethereum MainNet",
		},
	})
	if _, err := app.Run(); err != nil {
		println(err)
		os.Exit(1)
	}
}
