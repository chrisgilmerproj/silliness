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
		newColumn(instance),
	}
	m.focusedService = ec2Service
	m.services = []choice{
		newChoice(unselectedService, "Unselected"),
		newChoice(ec2Service, "EC2 Instances"),
		newChoice(ecsService, "ECS Tasks"),
	}

	if finalModel, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		log.Fatal(err)
	} else {
		instanceId := finalModel.(Model).resourceId
		if len(instanceId) > 0 {
			fmt.Println(finalModel.(Model).PrintCmd())
		}
		if finalModel.(Model).err != nil {
			log.Fatal(finalModel.(Model).err)
		}
	}
}
