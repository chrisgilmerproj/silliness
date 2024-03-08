package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

/*
 * aws ec2 describe-tags --filters='[{"Name":"resource-type","Values": ["instance"]},{"Name": "key","Values":["Name"]}]'
 */

type section int

const (
	tagKey section = iota
	tagValue
	instance
)

/* Styling */

var (
	docStyle = lipgloss.NewStyle().
			Margin(1, 2)
	columnStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.HiddenBorder())
	focusedStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62"))
	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241"))
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
	// return truncateString(strings.Join(t.values, ", "), 20)
	return fmt.Sprintf("Items: %d", len(t.values))
}

func (t Tag) Key() string {
	return t.name
}

func (t Tag) Values() []string {
	return t.values
}

/* Main Model */
type Model struct {
	focused  section
	data     GroupedKeyValueData
	lists    []list.Model
	err      error
	loaded   bool
	quitting bool
}

func New() *Model {
	return &Model{}
}

func (m *Model) SelectTagName() tea.Msg {
	selectedItem := m.lists[m.focused].SelectedItem()
	selectedTag := selectedItem.(Tag)
	switch selectedTag.section {
	case tagKey:
		newList := []list.Item{}
		for _, val := range selectedTag.Values() {
			instances := m.data[selectedTag.Key()][val]
			newList = append(newList, Tag{section: tagValue, name: val, values: instances})
		}
		m.lists[tagValue].SetItems(newList)
		m.lists[instance].SetItems([]list.Item{})
	case tagValue:
		newList := []list.Item{}
		for _, val := range selectedTag.Values() {
			newList = append(newList, Tag{section: tagValue, name: val, values: []string{}})
		}
		m.lists[instance].SetItems(newList)
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

func (m *Model) initLists(width, height int) {
	m.data = pullData()

	defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), width, height)
	defaultList.SetShowHelp(false)
	m.lists = []list.Model{defaultList, defaultList, defaultList}

	// Init Keys
	m.lists[tagKey].Title = "Key Names"
	var keyNameItems []list.Item
	for key, groupedValueData := range m.data {
		var values []string
		for val := range groupedValueData {
			values = append(values, val)
		}
		keyNameItems = append(keyNameItems, Tag{section: tagKey, name: key, values: values})
	}
	m.lists[tagKey].SetItems(keyNameItems)

	// Init Values as empty, fill this later
	m.lists[tagValue].Title = "Key Values"
	m.lists[tagValue].SetItems([]list.Item{})
	// Init Instances as empty, fill this later
	m.lists[instance].Title = "Instances"
	m.lists[instance].SetItems([]list.Item{})
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.loaded {
			_, v := docStyle.GetFrameSize()
			m.initLists(msg.Width, (msg.Height - v))
			m.loaded = true
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		case "left", "h":
			m.Prev()
		case "right", "l":
			m.Next()
		case "enter":
			return m, m.SelectTagName
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
	if m.loaded {
		tagKeyView := m.lists[tagKey].View()
		tagValueView := m.lists[tagValue].View()
		instanceView := m.lists[instance].View()
		switch m.focused {
		case tagValue:
			return docStyle.Render(lipgloss.JoinHorizontal(lipgloss.Left,
				columnStyle.Render(tagKeyView),
				focusedStyle.Render(tagValueView),
				columnStyle.Render(instanceView),
			))
		case instance:
			return docStyle.Render(lipgloss.JoinHorizontal(lipgloss.Left,
				columnStyle.Render(tagKeyView),
				columnStyle.Render(tagValueView),
				focusedStyle.Render(instanceView),
			))

		default:
			return docStyle.Render(lipgloss.JoinHorizontal(lipgloss.Left,
				focusedStyle.Render(tagKeyView),
				columnStyle.Render(tagValueView),
				columnStyle.Render(instanceView),
			))
		}
	} else {
		return docStyle.Render("loading ...")
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
