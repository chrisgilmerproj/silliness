package main

import (
	"encoding/json"
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
	commandStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("220"))
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
	switch t.section {
	case instance:
		return "instance"
	default:
		return fmt.Sprintf("Items: %d", len(t.values))
	}
}

func (t Tag) Key() string {
	return t.name
}

func (t Tag) Values() []string {
	return t.values
}

/* Main Model */
type Model struct {
	focused    section
	data       GroupedKeyValueData
	lists      []list.Model
	loaded     bool
	quitting   bool
	instanceId string
}

func New() *Model {
	return &Model{}
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
		newList := []list.Item{}
		for _, val := range selectedTag.Values() {
			instances := m.data[selectedTag.Key()][val]
			newList = append(newList, Tag{section: tagValue, name: val, values: instances})
		}
		m.lists[tagValue].SetItems(newList)
		m.lists[tagValue].ResetFilter()
		m.lists[instance].SetItems([]list.Item{})
		m.lists[instance].ResetFilter()
		m.Next()
	case tagValue:
		newList := []list.Item{}
		for _, val := range selectedTag.Values() {
			newList = append(newList, Tag{section: instance, name: val, values: []string{}})
		}
		m.lists[instance].SetItems(newList)
		m.lists[instance].ResetFilter()
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
			m.initLists(msg.Width, msg.Height*3/4)
			m.loaded = true
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		case "left":
			m.Prev()
		case "right":
			m.Next()
		case "enter":
			m.SelectListItem()
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
		doc := strings.Builder{}

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
		doc.WriteString(render)
		if len(m.instanceId) > 0 {
			doc.WriteString("\n\n" + commandStyle.Render(m.PrintCmd(m.instanceId, "")))
		}
		return docStyle.Render(doc.String())
	} else {
		return docStyle.Render("loading ...")
	}
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

func main() {
	m := New()
	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
