package main

import (
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
	chosenService  service
	services       []serviceChoice

	// List Management
	focusedColumn section
	columns       []column

	// Data
	data GroupedKeyValueData

	// Choices
	command *command

	// Other
	quitting bool
	err      error

	// Window
	width  int
	height int
}

func New() *Model {
	help := help.New()
	help.ShowAll = false
	return &Model{help: help, focusedColumn: tagKey}
}

func (m *Model) SelectListItem() tea.Msg {
	selectedItem := m.columns[m.focusedColumn].list.SelectedItem()
	// Move back a column if no items can be selected
	if selectedItem == nil {
		m.PrevColumn()
		return nil
	}
	// Process the selected item
	selectedTag := selectedItem.(SectionItem)
	switch selectedTag.section {
	case tagKey:

		values := selectedTag.Values()
		sort.Strings(values)

		newList := []list.Item{}
		for _, val := range values {
			resources := m.data[selectedTag.Key()][val]
			newList = append(newList, SectionItem{section: tagValue, name: val, values: resources})
		}
		m.columns[tagValue].list.SetItems(newList)
		m.columns[tagValue].list.ResetFilter()
		m.columns[resource].list.SetItems([]list.Item{})
		m.columns[resource].list.ResetFilter()
		m.ResetCommand()
		m.NextColumn()
	case tagValue:

		values := selectedTag.Values()
		sort.Strings(values)

		newList := []list.Item{}
		for _, val := range values {
			newList = append(newList, SectionItem{section: resource, name: val, values: []string{}})
		}
		m.columns[resource].list.SetItems(newList)
		m.columns[resource].list.ResetFilter()
		m.ResetCommand()
		m.NextColumn()
	case resource:
		switch m.chosenService {
		case ec2Service:
			m.command.resource = &ec2Choice{
				tag:        m.columns[tagKey].list.SelectedItem().(SectionItem).name,
				key:        m.columns[tagValue].list.SelectedItem().(SectionItem).name,
				instanceId: selectedTag.Key(),
			}
		case ecsService:
			m.command.resource = &ecsChoice{
				cluster:       m.columns[tagKey].list.SelectedItem().(SectionItem).name,
				containerName: m.columns[tagValue].list.SelectedItem().(SectionItem).name,
				taskId:        selectedTag.Key(),
			}
		}
	}
	return nil
}

func (m *Model) NextColumn() {
	m.columns[m.focusedColumn].Blur()
	m.focusedColumn = m.focusedColumn.getNext()
	m.columns[m.focusedColumn].Focus()
}

func (m *Model) PrevColumn() {
	m.columns[m.focusedColumn].Blur()
	m.focusedColumn = m.focusedColumn.getPrev()
	m.columns[m.focusedColumn].Focus()
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

func (m *Model) ResetCommand() {
	m.command = &command{}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		var cmd tea.Cmd
		var cmds []tea.Cmd
		m.width = msg.Width
		m.height = msg.Height
		m.help.Width = m.width - margin
		for i := 0; i < len(m.columns); i++ {
			var res tea.Model
			res, cmd = m.columns[i].Update(msg)
			m.columns[i] = res.(column)
			cmds = append(cmds, cmd)
		}
		if m.command != nil {
			_, cmd := m.command.Update(msg)
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
			if !m.columns[m.focusedColumn].list.SettingFilter() {
				switch {
				case key.Matches(msg, keys.Quit):
					m.quitting = true
					return m, tea.Quit
				case key.Matches(msg, keys.Left):
					m.PrevColumn()
				case key.Matches(msg, keys.Right):
					m.NextColumn()
				case key.Matches(msg, keys.Enter):
					m.SelectListItem()
				case key.Matches(msg, keys.Update):
					for _, c := range m.columns {
						c.list.SetItems([]list.Item{})
						c.list.ResetFilter()
					}
					m.initLists()
				case key.Matches(msg, keys.Run):
					return m, execCommand(m.command.resource.SliceCmd())
				case key.Matches(msg, keys.Switch):
					m.NextService()
					m.chosenService = m.focusedService
					for _, c := range m.columns {
						c.list.SetItems([]list.Item{})
						c.list.ResetFilter()
					}
					m.command = &command{}
					m.initLists()
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
	res, cmd := m.columns[m.focusedColumn].Update(msg)
	if _, ok := res.(column); ok {
		m.columns[m.focusedColumn] = res.(column)
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
		return lipgloss.Place(
			m.width,
			m.height,
			lipgloss.Center,
			lipgloss.Center,
			docStyle.Render(
				lipgloss.JoinVertical(
					lipgloss.Center,
					dialogBoxStyle.Render(ui),
				),
			),
		)
	}

	if len(m.data) == 0 {
		s := spinner.New()
		s.Spinner = spinner.Dot
		s.Style = spinnerStyle
		return docStyle.Render(fmt.Sprintf("%s loading from AWS ...", s.View()))
	}

	service := chosenServiceStyle.Render(m.services[m.chosenService].name)
	columns := lipgloss.JoinHorizontal(lipgloss.Left,
		m.columns[tagKey].View(),
		m.columns[tagValue].View(),
		m.columns[resource].View(),
	)
	cmdBlock := "\n"
	if m.command != nil {
		cmdBlock = m.command.View()
	}

	return docStyle.Render(
		lipgloss.JoinVertical(lipgloss.Center, service, columns, cmdBlock, m.help.View(keys)),
	)
}
