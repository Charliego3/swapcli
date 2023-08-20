package main

import "github.com/charmbracelet/lipgloss"

var (
	boldStyle = lipgloss.NewStyle().Bold(true)
	padding   = lipgloss.NewStyle().PaddingLeft(2)

	blue = boldStyle.Copy().Padding(1, 2).Foreground(lipgloss.AdaptiveColor{
		Light: "#0020b9",
		Dark:  "#0084e1",
	})

	optionStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.AdaptiveColor{
		Light: "#57007c",
		Dark:  "#ab00ff",
	})
	green = boldStyle.Copy().PaddingLeft(2).Foreground(lipgloss.AdaptiveColor{
		Light: "#005a2c",
		Dark:  "#15b900",
	})

	red = boldStyle.Copy().Padding(1, 2).Foreground(lipgloss.AdaptiveColor{
		Light: "#870004",
		Dark:  "#cd0004",
	})

	grey = boldStyle.Copy().Foreground(lipgloss.AdaptiveColor{
		Light: "#606269",
		Dark:  "#91949b",
	})
)
