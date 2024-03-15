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
	healthState := c.resource.HealthState()
	colors := map[string]string{
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
	healthColor := colors[healthState]
	return commandStyle.
		Width(width).
		BorderForeground(lipgloss.Color(healthColor)).
		Render(cmdString)
}
