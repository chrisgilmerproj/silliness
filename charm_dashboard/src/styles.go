package main

import "github.com/charmbracelet/lipgloss"

/* Styling */

var (
	docStyle = lipgloss.NewStyle().
			Margin(1, 2)
	columnStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.HiddenBorder())
	focusedStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62"))
	commandStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("220"))
	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241"))
)
