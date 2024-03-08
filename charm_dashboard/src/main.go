package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

/*
 * aws ec2 describe-tags --filters='[{"Name":"resource-type","Values": ["instance"]},{"Name": "key","Values":["Name"]}]'
 */

const divisor = 4

type section int

const (
	tagKey section = iota
	tagValue
)

/* Custom Tag */
type Tag struct {
	section section
	name    string
	values  []string
}

func (t Tag) FilterValue() string {
	return t.name
}

func (t Tag) Title() string {
	return t.name
}

func (t Tag) Description() string {
	return truncateString(strings.Join(t.values, ", "), 20)
}

/* Main Model */
type Model struct {
	focused section
	lists   []list.Model
	err     error
	loaded  bool
}

func New() *Model {
	return &Model{}
}

func (m *Model) initLists(width, height int) {
	defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), width/divisor-2, height)
	defaultList.SetShowHelp(false)
	m.lists = []list.Model{defaultList, defaultList}
	// Init Keys
	m.lists[tagKey].Title = "Key Names"
	m.lists[tagKey].SetItems([]list.Item{
		Tag{section: tagKey, name: "Name", values: []string{"admin", "deploy"}},
		Tag{section: tagKey, name: "ManagedBy", values: []string{"terraform", "cloudformation"}},
	})
	// Init Values as empty, fill this later
	m.lists[tagValue].Title = "Key Values"
	m.lists[tagValue].SetItems([]list.Item{

		Tag{section: tagKey, name: "ManagedBy", values: []string{"terraform", "cloudformation"}},
	})
}

// TODO: This is where we call aws tags
func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.loaded {
			m.initLists(msg.Width, msg.Height)
			m.loaded = true
		}
	}
	var cmd tea.Cmd
	m.lists[m.focused], cmd = m.lists[m.focused].Update(msg)
	return m, cmd
}

func (m Model) View() string {
	if m.loaded {
		tagKeyView := m.lists[tagKey].View()
		tagValueView := m.lists[tagValue].View()
		return lipgloss.JoinHorizontal(lipgloss.Left,
			tagKeyView,
			tagValueView,
		)
	} else {
		return "loading ..."
	}

}

func main() {
	m := New()
	//if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
