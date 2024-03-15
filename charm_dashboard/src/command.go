package main

import tea "github.com/charmbracelet/bubbletea"

type command struct {
	resource resourceChoice
	height   int
	width    int
}

func (c command) Init() tea.Cmd {
	return nil
}

// Update handles all the I/O for columns.
func (c command) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		c.setSize(msg.Width, msg.Height)
	}
	return c, cmd
}

func (c command) View() string {
	if c.resource == nil {
		return "\n"
	}
	return commandStyle.Render(c.resource.CmdToString())
}

func (c *command) setSize(width, height int) {
	c.width = width
	c.height = height
}
