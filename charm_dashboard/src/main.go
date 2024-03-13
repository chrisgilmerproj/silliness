package main

import (
	"fmt"
	"log"

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

func (s section) getNext() section {
	if s == instance {
		return tagKey
	}
	return s + 1
}

func (s section) getPrev() section {
	if s == tagKey {
		return instance
	}
	return s - 1
}

const (
	tagKey section = iota
	tagValue
	instance
)

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatalf("err: %w", err)
	}
	defer f.Close()

	m := New()

	m.cols = []column{
		newColumn(tagKey),
		newColumn(tagValue),
		newColumn(instance),
	}

	if finalModel, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		log.Fatalf("err: %w", err)
	} else {
		instanceId := finalModel.(Model).instanceId
		if len(instanceId) > 0 {
			fmt.Println(m.PrintCmd(instanceId, ""))
		}
		if finalModel.(Model).err != nil {
			log.Fatal(finalModel.(Model).err)
		}
	}
}
