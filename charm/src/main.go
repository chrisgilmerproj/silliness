package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const defaultListHeight = 14
const defaultListWidth = 20

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	docStyle          = lipgloss.NewStyle().Margin(1, 2)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type item string

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int  { return 1 }
func (d itemDelegate) Spacing() int { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	return nil
}
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

type model struct {
	list        list.Model
	quitting    bool
	groupedData GroupedKeyValueData
	key         string
	val         string
}

func (m model) Init() tea.Cmd {
	return func() tea.Msg {
		grouped := groupEC2Instances()
		ec2List := []list.Item{}
		for key, _ := range grouped {
			ec2List = append(ec2List, item(key))
		}
		return ec2List
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case []list.Item:
		m.list.SetItems([]list.Item(msg))
		return m, nil

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				if m.key == "" {
					m.key = string(i)
					ec2ValList := []list.Item{}
					for val := range m.groupedData[m.key] {
						ec2ValList = append(ec2ValList, item(val))
					}
					m.list = list.New(ec2ValList, itemDelegate{}, defaultListWidth, defaultListHeight)
					return m, nil
				} else if m.val == "" {
					m.val = string(i)
					ec2DataList := []list.Item{}
					for _, data := range m.groupedData[m.key][m.val] {
						ec2DataList = append(ec2DataList, item(data.InstanceId))
					}
					m.list = list.New(ec2DataList, itemDelegate{}, defaultListWidth, defaultListHeight)
					return m, nil
				}
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	// if len(m.list.Items()) == 0 {
	// 	return quitTextStyle.Render("Waiting for results from AWS.")
	// }
	if m.key != "" && m.val == "" {
		return quitTextStyle.Render(fmt.Sprintf("Looking at: %s, %v.", m.key))
	}
	if m.key != "" && m.val != "" {
		return quitTextStyle.Render(fmt.Sprintf("Looking at: %s %s.", m.key, m.val))
	}
	if m.quitting {
		return quitTextStyle.Render("Inspection complete.")
	}
	return docStyle.Render(m.list.View())
}

func main() {

	l := list.New([]list.Item{}, itemDelegate{}, defaultListWidth, defaultListHeight)
	l.Title = "EC2 Instance Tags"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	m := model{list: l}

	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
