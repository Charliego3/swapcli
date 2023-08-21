package main

import (
	"context"
	"fmt"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/ethereum/go-ethereum/ethclient"
	"os"
	"swapcli/contracts"
	"time"
)

type Options struct {
	chainId int
	client  *ethclient.Client
	token0  *entities.Token
	token1  *entities.Token
	erc200  *contracts.ERC20
	erc201  *contracts.ERC20
	market  string
}

const (
	BSC  = "BNB Smart Chain MainNet"
	Main = "Ethereum MainNet"
)

type KeyType int

const (
	KeyUp KeyType = 1 << iota
	KeyDown
	KeyHelp
	KeyQuit
	KeyEnter
)

type KeyMap struct {
	keys  map[KeyType]key.Binding
	short KeyType
	full  []KeyType
}

func (k KeyMap) ShortHelp() []key.Binding {
	return k.getBinding(k.short)
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		k.getBinding(k.full[0]),
		k.getBinding(k.full[1]),
	}
}

func (k KeyMap) getBinding(flag KeyType) []key.Binding {
	var keys []key.Binding
	if flag&KeyUp != 0 {
		keys = append(keys, k.keys[KeyUp])
	}
	if flag&KeyDown != 0 {
		keys = append(keys, k.keys[KeyDown])
	}
	if flag&KeyHelp != 0 {
		keys = append(keys, k.keys[KeyHelp])
	}
	if flag&KeyQuit != 0 {
		keys = append(keys, k.keys[KeyQuit])
	}
	if flag&KeyEnter != 0 {
		keys = append(keys, k.keys[KeyEnter])
	}
	return keys
}

var (
	opts = &Options{}

	KeyBindingQuit = key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q/esc/ctrl+c", "quit"),
	)
)

type ChainSelector struct {
	keyMap  KeyMap
	help    help.Model
	cursor  int
	choices []string
}

func (c ChainSelector) Init() tea.Cmd {
	return nil
}

func (c ChainSelector) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, c.keyMap.keys[KeyEnter]):
			chain := c.choices[c.cursor]
			var rawURL string
			if chain == BSC {
				opts.chainId = 56
				rawURL = "https://bscrpc.com"
			} else if chain == Main {
				opts.chainId = 1
				rawURL = "https://eth.drpc.org"
			}

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()
			client, err := ethclient.DialContext(ctx, rawURL)
			if err != nil {
				Println(red.Render(err.Error()))
				return c, tea.Quit
			}
			opts.client = client
			s := spinner.New(
				spinner.WithSpinner(spinner.Globe),
				spinner.WithStyle(optionStyle),
			)
			return NewTokenSelector(s), tea.Batch(tea.Println(
				green.Render("\nNetwork: "),
				grey.Render(c.choices[c.cursor]),
			), textinput.Blink, s.Tick)
		case key.Matches(msg, c.keyMap.keys[KeyDown]):
			c.cursor++
			if c.cursor > len(c.choices)-1 {
				c.cursor = 0
			}
		case key.Matches(msg, c.keyMap.keys[KeyUp]):
			c.cursor--
			if c.cursor < 0 {
				c.cursor = len(c.choices) - 1
			}
		case key.Matches(msg, c.keyMap.keys[KeyHelp]):
			c.help.ShowAll = !c.help.ShowAll
		case key.Matches(msg, c.keyMap.keys[KeyQuit]):
			return c, tea.Quit
		}
	}

	return c, nil
}

func (c ChainSelector) View() string {
	if opts.chainId > 0 {
		return ""
	}

	s := blue.Render("What's network do you want to choose?") + "\n"
	for i, choice := range c.choices {
		if c.cursor == i { //    T
			s += green.Render(fmt.Sprintf("\U000F012D %s", choice)) + "\n"
		} else {
			s += optionStyle.Render(fmt.Sprintf("  %s", choice)) + "\n"
		}
	}

	s += padding.Render("\n" + c.help.View(c.keyMap) + "\n")
	return s
}

func Println(v ...any) {
	go app.Println(v...)
	time.Sleep(time.Millisecond)
}

var (
	app = tea.NewProgram(ChainSelector{
		choices: []string{BSC, Main},
		help:    help.New(),
		keyMap:  NewSelectKeyMap(),
	})
)

func main() {
	if _, err := app.Run(); err != nil {
		println(err)
		os.Exit(1)
	}
}
