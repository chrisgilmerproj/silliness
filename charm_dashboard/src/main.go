package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
)

func main() {
	// Initialize the gum log
	logger := log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    false,
		ReportTimestamp: false,
		TimeFormat:      time.RFC3339,
		Prefix:          "ssmpicker ",
		Level:           log.InfoLevel,
	})

	f, err := tea.LogToFile("debug.log", "ssmpicker")
	if err != nil {
		logger.Error(err)
	}
	defer f.Close()

	// If the AWS Environment variables are not set, the program will not work
	if os.Getenv("AWS_ACCESS_KEY_ID") == "" || os.Getenv("AWS_SECRET_ACCESS_KEY") == "" || os.Getenv("AWS_SESSION_TOKEN") == "" {
		logger.Fatal("AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY and AWS_SESSION_TOKEN must be set")
	}

	// Configure AWS Clients
	var errEc2 error
	ec2Client, errEc2 = GetEC2Client()
	if errEc2 != nil {
		logger.Fatal(errEc2)
	}
	var errEcs error
	ecsClient, errEcs = GetECSClient()
	if errEcs != nil {
		logger.Fatal(errEcs)
	}

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
		logger.Fatal(err)
	} else {
		fm := finalModel.(Model)
		if fm.command != nil {
			fmt.Println(fm.command.resource.CmdToString())
		}
		if fm.err != nil {
			logger.Fatal(fm.err)
		}
	}
}
