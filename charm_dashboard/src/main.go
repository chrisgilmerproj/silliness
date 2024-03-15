package main

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

/*
 * Data from:
 * aws ec2 describe-tags --filters='[{"Name":"resource-type","Values": ["resource"]}]'
 *
 * Compare filtered result for Name to:
 * aws ec2 describe-tags --filters='[{"Name":"resource-type","Values": ["resource"]},{"Name": "key","Values":["Name"]}]'
 */

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

	m.cols = []column{
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
		if m.ec2Choice != nil {
			fmt.Println(finalModel.(Model).ec2Choice.CmdToString())
		} else if m.ecsChoice != nil {
			fmt.Println(finalModel.(Model).ecsChoice.CmdToString())
		}
		if finalModel.(Model).err != nil {
			log.Fatal(finalModel.(Model).err)
		}
	}
}
