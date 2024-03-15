package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type serviceChoice struct {
	focus   bool
	service service
	name    string
}

func (c *serviceChoice) Focus() {
	c.focus = true
}

func (c *serviceChoice) Blur() {
	c.focus = false
}

func (c *serviceChoice) Focused() bool {
	return c.focus
}

func newChoice(service service, name string) serviceChoice {
	var focus bool
	if service == ec2Service {
		focus = true
	}
	return serviceChoice{focus: focus, service: service, name: name}
}

func (c serviceChoice) Init() tea.Cmd {
	return nil
}

func (c serviceChoice) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	return c, cmd
}

func (c serviceChoice) View() string {
	return c.getStyle().Render(c.name)
}

func (c *serviceChoice) getStyle() lipgloss.Style {
	if c.Focused() {
		return serviceFocusedStyle
	}
	return serviceStyle
}
