package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

/* Main Model */
type Model struct {
	help       help.Model
	lists      []list.Model
	focused    section
	data       GroupedKeyValueData
	quitting   bool
	instanceId string
}

func New() *Model {
	help := help.New()
	help.ShowAll = true
	return &Model{help: help, focused: tagKey}
}

func (m *Model) SelectListItem() tea.Msg {
	selectedItem := m.lists[m.focused].SelectedItem()
	// Move back a column if no items can be selected
	if selectedItem == nil {
		m.Prev()
		return nil
	}
	// Process the selected item
	selectedTag := selectedItem.(Tag)
	switch selectedTag.section {
	case tagKey:

		values := selectedTag.Values()
		sort.Strings(values)

		newList := []list.Item{}
		for _, val := range values {
			instances := m.data[selectedTag.Key()][val]
			newList = append(newList, Tag{section: tagValue, name: val, values: instances})
		}
		m.lists[tagValue].SetItems(newList)
		m.lists[tagValue].ResetFilter()
		m.lists[instance].SetItems([]list.Item{})
		m.lists[instance].ResetFilter()
		m.instanceId = ""
		m.Next()
	case tagValue:

		values := selectedTag.Values()
		sort.Strings(values)

		newList := []list.Item{}
		for _, val := range values {
			newList = append(newList, Tag{section: instance, name: val, values: []string{}})
		}
		m.lists[instance].SetItems(newList)
		m.lists[instance].ResetFilter()
		m.instanceId = ""
		m.Next()
	case instance:
		m.instanceId = selectedTag.Key()
	}
	return nil
}

func (m *Model) Next() {
	if m.focused == instance {
		m.focused = tagKey
	} else {
		m.focused++
	}
}

func (m *Model) Prev() {
	if m.focused == tagKey {
		m.focused = instance
	} else {
		m.focused--
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if len(m.data) == 0 {
			width := msg.Width
			height := msg.Height * 3 / 4

			defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), width, height)
			defaultList.SetShowHelp(false)
			m.lists = []list.Model{defaultList, defaultList, defaultList}
			m.initLists()
		}
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Quit):
			m.quitting = true
			return m, tea.Quit
		case key.Matches(msg, keys.Left):
			m.Prev()
		case key.Matches(msg, keys.Right):
			m.Next()
		case key.Matches(msg, keys.Enter):
			m.SelectListItem()
		case key.Matches(msg, keys.Update):
			for _, l := range m.lists {
				l.SetItems([]list.Item{})
			}
			m.initLists()
		case key.Matches(msg, keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		}

	}
	var cmd tea.Cmd
	m.lists[m.focused], cmd = m.lists[m.focused].Update(msg)
	return m, cmd
}

func (m Model) View() string {
	if m.quitting {
		return docStyle.Render("")
	}

	if len(m.data) == 0 {
		return docStyle.Render("loading ...")
	}

	tagKeyView := m.lists[tagKey].View()
	tagValueView := m.lists[tagValue].View()
	instanceView := m.lists[instance].View()

	var render string
	switch m.focused {
	case tagValue:
		render = lipgloss.JoinHorizontal(lipgloss.Left,
			columnStyle.Render(tagKeyView),
			focusedStyle.Render(tagValueView),
			columnStyle.Render(instanceView),
		)
	case instance:
		render = lipgloss.JoinHorizontal(lipgloss.Left,
			columnStyle.Render(tagKeyView),
			columnStyle.Render(tagValueView),
			focusedStyle.Render(instanceView),
		)
	default:
		render = lipgloss.JoinHorizontal(lipgloss.Left,
			focusedStyle.Render(tagKeyView),
			columnStyle.Render(tagValueView),
			columnStyle.Render(instanceView),
		)
	}

	cmdBlock := "\n"
	if len(m.instanceId) > 0 {
		cmdBlock = commandStyle.Render(m.PrintCmd(m.instanceId, ""))
	}

	return docStyle.Render(
		lipgloss.JoinVertical(lipgloss.Left, render, cmdBlock, m.help.View(keys)),
	)
}

func (m Model) PrintCmd(instanceId, portForwarding string) string {
	command := []string{
		"aws",
		"ssm",
		"start_session",
		"--target",
		instanceId,
	}
	if len(portForwarding) > 0 {
		portFromTo := strings.Split(portForwarding, ":")
		ports := map[string][]string{
			"portNumber":      {portFromTo[0]},
			"localPortNumber": {portFromTo[1]},
		}
		compactPorts, _ := json.Marshal(ports)
		command = append(command, "--document-name AWS-StartPortForwardingSession")
		command = append(command, fmt.Sprintf("--parameters '%s'", compactPorts))
	}
	return strings.Join(command, " ")
}
