package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const APPEND = -1

type column struct {
	focus   bool
	section section
	list    list.Model
	height  int
	width   int
}

func (c *column) Focus() {
	c.focus = true
}

func (c *column) Blur() {
	c.focus = false
}

func (c *column) Focused() bool {
	return c.focus
}

func newColumn(section section) column {
	var focus bool
	if section == tagKey {
		focus = true
	}
	defaultDelegate := list.NewDefaultDelegate()
	defaultDelegate.SetSpacing(0)
	defaultDelegate.ShowDescription = false
	defaultList := list.New([]list.Item{}, defaultDelegate, 0, 0)
	defaultList.SetShowHelp(false)
	defaultList.DisableQuitKeybindings()
	return column{focus: focus, section: section, list: defaultList}
}

// Init does initial setup for the column.
func (c column) Init() tea.Cmd {
	return nil
}

// Update handles all the I/O for columns.
func (c column) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		c.setSize(msg.Width, msg.Height)
		c.list.SetSize(msg.Width/margin, msg.Height/2)
	}
	c.list, cmd = c.list.Update(msg)
	return c, cmd
}

func (c column) View() string {
	return c.getStyle().Render(c.list.View())
}

func (c *column) setSize(width, height int) {
	c.width = width / margin
	c.height = height / 2
}

func (c *column) getStyle() lipgloss.Style {
	if c.Focused() {
		return focusedStyle.
			Height(c.height).
			Width(c.width)
	}
	return columnStyle.
		Height(c.height).
		Width(c.width)
}
