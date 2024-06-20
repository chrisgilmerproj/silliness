package main

import (
	"math"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type command struct {
	resource resourceChoice
}

func (c command) Init() tea.Cmd {
	return nil
}

// Update handles all the I/O for columns.
func (c command) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	return c, cmd
}

func (c command) View() string {
	if c.resource == nil {
		return ""
	}
	maxWidth := 80
	// Don't let wrapping split words
	cmdString := splitLine(c.resource.SliceCmd(), maxWidth)
	// The +4 accounts for padding on the left/right of the box
	width := int(math.Min(float64(len(cmdString))+4, float64(maxWidth)))
	healthState, _ := c.resource.HealthState()
	healthColor := colorHealthMap[healthState]
	return commandStyle.
		Width(width).
		BorderForeground(lipgloss.Color(healthColor)).
		Render(cmdString)
}
