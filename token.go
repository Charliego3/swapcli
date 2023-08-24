package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/ethereum/go-ethereum/common"
	"strconv"
	"strings"
	"swapcli/contracts"
)

type TokenSelector struct {
	help     help.Model
	keyMap   KeyMap
	input    textinput.Model
	spinner  spinner.Model
	quitting bool
	loading  bool
	token0o  bool
	prompt   string
	err      error
}

func NewTokenSelector(s spinner.Model) *TokenSelector {
	return &TokenSelector{
		help:    help.New(),
		input:   NewTextInput("please enter the contract address"),
		spinner: s,
		keyMap: KeyMap{
			keys: map[KeyType]key.Binding{
				KeyEnter: NewEnterBinding("confirm"),
				KeyQuit:  KeyBindingQuit,
			},
			short: KeyEnter | KeyQuit,
		},
	}
}

func (t *TokenSelector) Init() tea.Cmd {
	return nil
}

func (t *TokenSelector) getPrompt() string {
	if opts.token0 == nil {
		return "Pricing Currency"
	}
	return "Base Currency"
}

func (t *TokenSelector) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if t.quitting {
		return t, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, t.keyMap.keys[KeyEnter]):
			if t.loading {
				break
			}

			t.loading = true
			delete(t.keyMap.keys, KeyEnter)
			go func() {
				defer func() {
					t.loading = false
					t.keyMap.keys[KeyEnter] = NewEnterBinding("confirm")
				}()
				addressHex := t.input.Value()
				if !common.IsHexAddress(addressHex) {
					t.err = fmt.Errorf("contract address is invalid: %q", addressHex)
					return
				}

				var erc20 *contracts.ERC20
				address := common.HexToAddress(addressHex)
				erc20, t.err = contracts.NewERC20(address, opts.client)
				if t.err != nil {
					return
				}

				var name, symbol string
				name, t.err = erc20.Name(nil)
				if t.err != nil {
					return
				}

				symbol, t.err = erc20.Symbol(nil)
				if t.err != nil {
					return
				}

				var decimals uint8
				decimals, t.err = erc20.Decimals(nil)
				if t.err != nil {
					return
				}

				token := entities.NewToken(uint(opts.chainId), address, uint(decimals), symbol, name)
				t.prompt = strings.Replace(t.getPrompt(), " ", "", 1)
				if opts.token0 == nil {
					opts.erc200 = erc20
					opts.token0 = token
				} else {
					opts.erc201 = erc20
					opts.token1 = token
				}
			}()
		case key.Matches(msg, t.keyMap.keys[KeyQuit]):
			t.quitting = true
			return t, tea.Quit
		}
	}

	if t.err != nil {
		return t, tea.Quit
	}

	if opts.token1 != nil {
		t.quitting = true
		info := t.outputToken(t.prompt, opts.token1)
		return NewMarketConfirm(), tea.Batch(tea.Println(info...), textinput.Blink)
	}

	var sCmd, iCmd tea.Cmd
	t.spinner, sCmd = t.spinner.Update(msg)
	t.input, iCmd = t.input.Update(msg)
	if opts.token0 != nil && !t.token0o {
		t.token0o = true
		t.input.Reset()
		info := t.outputToken(t.prompt, opts.token0)
		return t, tea.Batch(sCmd, iCmd, t.input.Focus(), tea.Println(info...))
	}
	return t, tea.Batch(sCmd, iCmd)
}

func (t *TokenSelector) outputToken(prompt string, token *entities.Token) []any {
	return []any{
		green.Render(prompt + ":"),
		optionStyle.Render("\n  Address: "),
		grey.Render(token.Address.String()),
		optionStyle.Render("\n  Name: "),
		grey.Render(token.Name()),
		optionStyle.Render("\n  Symbol: "),
		grey.Render(token.Symbol()),
		optionStyle.Render("\n  Decimals: "),
		grey.Render(strconv.Itoa(int(token.Decimals()))),
	}
}

func (t *TokenSelector) View() string {
	if t.err != nil {
		return red.Render(t.err.Error()) + "\n"
	}

	if t.quitting {
		return "\n"
	}

	var s string
	if t.loading {
		s = "\n" + t.spinner.View() + grey.Render(" loading token...")
	} else {
		s = blue.Render("What's the contract address of the "+t.getPrompt()+"?") + "\n"
		s += t.input.View()
	}
	s += padding.Render("\n\n" + t.help.View(t.keyMap) + "\n")
	return s
}
