package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

/*
 * Data from:
 * aws ec2 describe-tags --filters='[{"Name":"resource-type","Values": ["instance"]}]'
 *
 * Compare filtered result for Name to:
 * aws ec2 describe-tags --filters='[{"Name":"resource-type","Values": ["instance"]},{"Name": "key","Values":["Name"]}]'
 */

type section int

const (
	tagKey section = iota
	tagValue
	instance
)

func main() {
	m := New()
	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
