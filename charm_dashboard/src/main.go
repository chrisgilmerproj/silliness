package main

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	f, err := tea.LogToFile("debug.log", "ssmpicker")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Configure AWS Clients
	ec2Client = GetEC2Client()
	ecsClient = GetECSClient()

	m := New()

	m.columns = []column{
		newColumn(tagKey),
		newColumn(tagValue),
		newColumn(resource),
	}
	m.focusedService = ec2Service
	m.services = []serviceChoice{
		newChoice(unselectedService, "Unselected"),
		newChoice(ec2Service, "EC2 resources"),
		newChoice(ecsService, "ECS Tasks"),
	}

	if finalModel, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		log.Fatal(err)
	} else {
		fm := finalModel.(Model)
		if fm.command != nil {
			fmt.Println(fm.command.resource.CmdToString())
		}
		if fm.err != nil {
			log.Fatal(fm.err)
		}
	}
}
