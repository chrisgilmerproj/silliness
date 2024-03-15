package main

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type section int

func (s section) getNext() section {
	if s == resource {
		return tagKey
	}
	return s + 1
}

func (s section) getPrev() section {
	if s == tagKey {
		return resource
	}
	return s - 1
}

const (
	tagKey section = iota
	tagValue
	resource
)

type service int

const (
	unselectedService service = iota
	ec2Service
	ecsService
)

func (s service) getNext() service {
	if s == ecsService {
		return ec2Service
	}
	return s + 1
}

func (s service) getPrev() service {
	if s == ec2Service {
		return ecsService
	}
	return s - 1
}

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	m := New()

	m.columns = []column{
		newColumn(tagKey),
		newColumn(tagValue),
		newColumn(resource),
	}
	m.focusedService = ec2Service
	m.services = []choice{
		newChoice(unselectedService, "Unselected"),
		newChoice(ec2Service, "EC2 resources"),
		newChoice(ecsService, "ECS Tasks"),
	}

	if finalModel, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		log.Fatal(err)
	} else {
		fm := finalModel.(Model)
		fmt.Println(fm.command.resource.CmdToString())
		if fm.err != nil {
			log.Fatal(fm.err)
		}
	}
}
