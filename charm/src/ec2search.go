package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	flagKey            = "key"
	flagVal            = "val"
	flagHealthStatus   = "health-status"
	flagPortForwarding = "port-forwarding"
	flagRun            = "run"

	defaultListHeight = 14
	defaultListWidth  = 20
)

var (
	choicesHealthStatus = []string{"pending", "running", "shutting-down", "terminated"}
)

// Initialize the flags for the subcommand
func initEc2SearchFlags(flag *pflag.FlagSet) {
	flag.String(flagKey, "", "The key to search on")
	flag.String(flagVal, "", "The val to search on")
	flag.String(flagHealthStatus, "", "The health status to filter on [todo: defaults]")
	flag.String(flagPortForwarding, "", "Enable port forwarding in the format 'from:to'")
	flag.Bool(flagRun, false, "Execute the command")
}

// Check that the configuration for the subcommand is correct
func checkEc2SearchConfig(v *viper.Viper, args []string) error {
	healthStatusFound := false
	healthStatus := v.GetString(flagHealthStatus)
	if len(healthStatus) > 0 {
		for _, status := range choicesHealthStatus {
			if healthStatus == status {
				healthStatusFound = true
			}
		}
		if !healthStatusFound {
			return fmt.Errorf("Health Status must be one of %v, you provided: %q", choicesHealthStatus, healthStatus)
		}
	}

	portForwarding := v.GetString(flagPortForwarding)
	if len(portForwarding) > 0 {
		if strings.Contains(portForwarding, ":") {
			fromToPort := strings.Split(portForwarding, ":")
			if len(fromToPort) != 2 {
				return fmt.Errorf("Port Forwarding format must be 'from:to', you provided: %q", portForwarding)
			}
			if _, err := strconv.Atoi(fromToPort[0]); err != nil {
				return fmt.Errorf("Port Forwarding 'from' port must be an integer, you provided: %q", fromToPort[0])
			}
			if _, err := strconv.Atoi(fromToPort[1]); err != nil {
				return fmt.Errorf("Port Forwarding 'to' port must be an integer, you provided: %q", fromToPort[1])
			}
		} else {
			return fmt.Errorf("Port Forwarding format must be 'from:to', you provided: %q", portForwarding)
		}
	}
	return nil
}

// The ec2 search subcommand
func ec2Search(cmd *cobra.Command, args []string) error {
	v, errViper := initViper(cmd)
	if errViper != nil {
		return fmt.Errorf("error initializing viper: %w", errViper)
	}

	if len(args) > 0 {
		return cmd.Usage()
	}

	if errConfig := checkEc2SearchConfig(v, args); errConfig != nil {
		return errConfig
	}

	// Get the flag values

	l := list.New([]list.Item{}, itemDelegate{}, defaultListWidth, defaultListHeight)
	l.Title = "EC2 Instance Tags"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	m := model{
		list:           l,
		key:            v.GetString(flagKey),
		val:            v.GetString(flagVal),
		healthStatus:   v.GetString(flagHealthStatus),
		portForwarding: v.GetString(flagPortForwarding),
		runCommand:     v.GetBool(flagRun),
	}
	m.groupedData = groupEC2Instances()

	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	return nil
}

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
	list           list.Model
	quitting       bool
	groupedData    GroupedKeyValueData
	key            string
	val            string
	instanceId     string
	healthStatus   string
	portForwarding string
	runCommand     bool
}

func (m model) Init() tea.Cmd {
	return m.UpdateList()
}

func (m model) UpdateList() tea.Cmd {
	return func() tea.Msg {
		keyList := []string{}
		if m.key != "" && m.val == "" {
			// When the key is chosen but not the val
			for val := range m.groupedData[m.key] {
				keyList = append(keyList, val)
			}
			sort.Strings(keyList)
			ec2ValList := []list.Item{}
			for _, val := range keyList {
				ec2ValList = append(ec2ValList, item(val))
			}
			return ec2ValList
		} else if m.key != "" && m.val != "" {
			// When both the key and val are chosen
			ec2DataList := []list.Item{}
			if groupedData, exists := m.groupedData[m.key][m.val]; exists {
				for _, data := range groupedData {
					ec2DataList = append(ec2DataList, item(data.InstanceId))
				}
			} else {
				for val := range m.groupedData[m.key] {
					if strings.Contains(val, m.val) {
						for _, data := range m.groupedData[m.key][val] {
							ec2DataList = append(ec2DataList, item(data.InstanceId))
						}
					}
				}
			}
			return ec2DataList
		}

		// When neither the key or val is chosen
		ec2List := []list.Item{}
		for key := range m.groupedData {
			keyList = append(keyList, key)
		}
		sort.Strings(keyList)
		for _, key := range keyList {
			ec2List = append(ec2List, item(key))
		}
		return ec2List
	}
}

func (m model) PrintCmd() string {
	command := []string{
		"aws",
		"ssm",
		"start_session",
		"--target",
		m.instanceId,
	}
	if len(m.portForwarding) > 0 {
		portFromTo := strings.Split(m.portForwarding, ":")
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
				if len(m.key) == 0 {
					m.key = string(i)
					return m, m.UpdateList()
				} else if len(m.val) == 0 {
					m.val = string(i)
					return m, m.UpdateList()
				} else {
					m.instanceId = string(i)
					m.quitting = true
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
	if len(m.list.Items()) == 0 {
		return quitTextStyle.Render("Waiting for results from AWS.")
	}
	// Render which values have been selected into the view
	// if m.key != "" && m.val == "" {
	// 	return quitTextStyle.Render(fmt.Sprintf("Looking at: %s.", m.key))
	// }
	// if m.key != "" && m.val != "" {
	// 	return quitTextStyle.Render(fmt.Sprintf("Looking at: %s %s.", m.key, m.val))
	// }
	if m.quitting {
		if len(m.key) != 0 && len(m.val) != 0 {
			return quitTextStyle.Render(m.PrintCmd())
		}
		return quitTextStyle.Render("Done.")
	}
	return docStyle.Render(m.list.View())
}
