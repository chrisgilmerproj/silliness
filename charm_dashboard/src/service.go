package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type choice struct {
	focus   bool
	service service
	name    string
}

func (c *choice) Focus() {
	c.focus = true
}

func (c *choice) Blur() {
	c.focus = false
}

func (c *choice) Focused() bool {
	return c.focus
}

func newChoice(service service, name string) choice {
	var focus bool
	if service == ec2Service {
		focus = true
	}
	return choice{focus: focus, service: service, name: name}
}

func (c choice) Init() tea.Cmd {
	return nil
}

func (c choice) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	return c, cmd
}

func (c choice) View() string {
	return c.getStyle().Render(c.name)
}

func (c *choice) getStyle() lipgloss.Style {
	if c.Focused() {
		return serviceFocusedStyle
	}
	return serviceStyle
}
