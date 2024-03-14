package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const margin = 4

type commandFinishedMsg struct{ err error }

func execCommand(command []string) tea.Cmd {
	env := os.Environ()
	// Filter environment variables that start with "AWS"
	awsEnv := make([]string, 0)
	for _, v := range env {
		if strings.HasPrefix(v, "AWS") || strings.HasPrefix(v, "PATH") {
			awsEnv = append(awsEnv, v)
		}
	}
	c := exec.Command(command[0], command[1:]...)
	c.Env = awsEnv
	return tea.ExecProcess(c, func(err error) tea.Msg {
		return commandFinishedMsg{err}
	})
}

/* Main Model */
type Model struct {
	// Help Management
	help help.Model

	// Service Management
	focusedService service
	services       []choice
	chosenService  service

	// List Management
	focused    section
	cols       []column
	instanceId string
	data       GroupedKeyValueData

	// Other
	quitting bool
	err      error
}

func New() *Model {
	help := help.New()
	help.ShowAll = true
	return &Model{help: help, focused: tagKey}
}

func (m *Model) SelectListItem() tea.Msg {
	selectedItem := m.cols[m.focused].list.SelectedItem()
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
		m.cols[tagValue].list.SetItems(newList)
		m.cols[tagValue].list.ResetFilter()
		m.cols[instance].list.SetItems([]list.Item{})
		m.cols[instance].list.ResetFilter()
		m.instanceId = ""
		m.Next()
	case tagValue:

		values := selectedTag.Values()
		sort.Strings(values)

		newList := []list.Item{}
		for _, val := range values {
			newList = append(newList, Tag{section: instance, name: val, values: []string{}})
		}
		m.cols[instance].list.SetItems(newList)
		m.cols[instance].list.ResetFilter()
		m.instanceId = ""
		m.Next()
	case instance:
		m.instanceId = selectedTag.Key()
	}
	return nil
}

func (m *Model) Next() {
	m.cols[m.focused].Blur()
	m.focused = m.focused.getNext()
	m.cols[m.focused].Focus()
}

func (m *Model) Prev() {
	m.cols[m.focused].Blur()
	m.focused = m.focused.getPrev()
	m.cols[m.focused].Focus()
}

func (m *Model) SelectService() tea.Msg {
	return nil
}

func (m *Model) NextService() {
	m.services[m.focusedService].Blur()
	m.focusedService = m.focusedService.getNext()
	m.services[m.focusedService].Focus()
}

func (m *Model) PrevService() {
	m.services[m.focusedService].Blur()
	m.focusedService = m.focusedService.getPrev()
	m.services[m.focusedService].Focus()
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		var cmd tea.Cmd
		var cmds []tea.Cmd
		m.help.Width = msg.Width - margin
		for i := 0; i < len(m.cols); i++ {
			var res tea.Model
			res, cmd = m.cols[i].Update(msg)
			m.cols[i] = res.(column)
			cmds = append(cmds, cmd)
		}
		return m, tea.Batch(cmds...)
	case tea.KeyMsg:
		// Until a service is selected use a different keyset
		if m.chosenService == 0 {
			switch {
			case key.Matches(msg, keys.Quit):
				m.quitting = true
				return m, tea.Quit
			case key.Matches(msg, keys.Left):
				m.PrevService()
			case key.Matches(msg, keys.Right):
				m.NextService()
			case key.Matches(msg, keys.Enter):
				m.chosenService = m.focusedService
				m.initLists()
			}
		} else {
			// If the filter is in use then do capture keys
			if !m.cols[m.focused].list.SettingFilter() {
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
					for _, c := range m.cols {
						c.list.SetItems([]list.Item{})
						c.list.ResetFilter()
					}
					m.initLists()
				case key.Matches(msg, keys.Run):
					if len(m.instanceId) > 0 {
						return m, execCommand(m.SliceCmd(m.instanceId, ""))
					}
				case key.Matches(msg, keys.Help):
					m.help.ShowAll = !m.help.ShowAll
				}
			}
		}
	case commandFinishedMsg:
		if msg.err != nil {
			m.err = msg.err
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	res, cmd := m.cols[m.focused].Update(msg)
	if _, ok := res.(column); ok {
		m.cols[m.focused] = res.(column)
	} else {
		return res, cmd
	}
	return m, cmd
}

func (m Model) View() string {
	if m.quitting {
		return docStyle.Render("")
	}

	if m.chosenService == 0 {
		question := servicePickerStyle.Render("Which service?")
		buttons := lipgloss.JoinHorizontal(
			lipgloss.Left,
			m.services[ec2Service].View(),
			m.services[ecsService].View(),
		)
		ui := lipgloss.JoinVertical(
			lipgloss.Center,
			question,
			buttons,
		)
		return lipgloss.Place(50, 9,
			lipgloss.Center, lipgloss.Center,
			dialogBoxStyle.Render(ui),
		)
	}

	if len(m.data) == 0 {
		s := spinner.New()
		s.Spinner = spinner.Dot
		s.Style = spinnerStyle
		return docStyle.Render(fmt.Sprintf("%s loading from AWS ...", s.View()))
	}

	var render string
	render = lipgloss.JoinHorizontal(lipgloss.Left,
		m.cols[tagKey].View(),
		m.cols[tagValue].View(),
		m.cols[instance].View(),
	)

	cmdBlock := "\n"
	if len(m.instanceId) > 0 {
		cmdBlock = commandStyle.Render(m.PrintCmd(m.instanceId, ""))
	}

	return docStyle.Render(
		lipgloss.JoinVertical(lipgloss.Left, render, cmdBlock, m.help.View(keys)),
	)
}

func (m Model) SliceCmd(instanceId, portForwarding string) []string {
	command := []string{}
	if m.chosenService == ec2Service {
		command = []string{
			"aws",
			"ssm",
			"start-session",
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
	} else if m.chosenService == ecsService {
		clusterId := m.cols[tagKey].list.SelectedItem().(Tag).name
		containerId := m.cols[tagValue].list.SelectedItem().(Tag).name
		command = []string{
			"aws",
			"ecs",
			"execute-command",
			"--cluster",
			clusterId,
			"--container",
			containerId,
			"--task",
			instanceId,
			"--interactive",
			"--command",
			"/bin/bash",
		}
	}
	return command
}

func (m Model) PrintCmd(instanceId, portForwarding string) string {
	command := m.SliceCmd(instanceId, portForwarding)
	return strings.Join(command, " ")
}
