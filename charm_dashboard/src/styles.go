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
			Border(lipgloss.RoundedBorder())
	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241"))
	spinnerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("205"))
	chosenServiceStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("220"))
	servicePickerStyle = lipgloss.NewStyle().
				Width(40).
				Align(lipgloss.Center)
	serviceStyle = lipgloss.NewStyle().
			Padding(1, 2)
	serviceFocusedStyle = serviceStyle.
				Copy().
				Foreground(lipgloss.Color("#FFF7DB")).
				Background(lipgloss.Color("#F25D94")).
				Underline(true)
	dialogBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#874BFD")).
			Padding(1, 0).
			BorderTop(true).
			BorderLeft(true).
			BorderRight(true).
			BorderBottom(true)
	colorHealthMap = map[string]string{
		// Unable to return data
		"": "220", // yellow
		// EC2
		"pending":       "220", // yellow
		"running":       "28",  // green
		"shutting-down": "163", // magenta
		"terminated":    "124", // red
		// ECS
		"healthy":   "28",  // green
		"unknown":   "220", // yellow
		"unhealthy": "124", // red
	}
)
