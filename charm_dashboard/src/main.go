package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/spf13/pflag"
)

var VERSION = "0.1.0"

func main() {

	// Get the command-line arguments
	var portForwardingFlag string
	var versionFlag bool
	var helpFlag bool
	pflag.StringVar(&portForwardingFlag, "port-forwarding", "", "Port forwarding in the format 'portNumber:localPortNumber'")
	pflag.BoolVar(&versionFlag, "version", false, "Show version and exit")
	pflag.BoolVarP(&helpFlag, "help", "h", false, "Show help and exit")
	pflag.Parse()

	if helpFlag {
		fmt.Println("ssmpicker is a tool to help you pick the right AWS SSM command to run on your resources.")
		fmt.Println("Usage: ssmpicker [ec2 ecs] [--version] [--help]")
		fmt.Println("Options:")
		pflag.PrintDefaults()
		return
	}

	// Iterate over the arguments to check for "--version"
	if versionFlag {
		fmt.Printf("v%s", VERSION)
		return
	}

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

	m := New()

	m.columns = []column{
		newColumn(tagKey),
		newColumn(tagValue),
		newColumn(resource),
	}
	m.services = []serviceChoice{
		newChoice(unselectedService, "Unselected"),
		newChoice(ec2Service, "EC2 resources"),
		newChoice(ecsService, "ECS Tasks"),
	}

	services := pflag.Args()
	allowedServices := []string{"ec2", "ecs"}
	if len(services) == 0 {
		services = allowedServices
	}

	// check that all services are allowed
	for _, service := range services {
		switch service {
		case "ec2":
			logger.Info("Service allowed", "service", service)
			// Configure AWS Clients
			var errEc2 error
			ec2Client, errEc2 = GetEC2Client()
			if errEc2 != nil {
				logger.Fatal(errEc2)
			}
			m.focusedService = ec2Service

			if len(portForwardingFlag) > 0 {
				logger.Info("Port forwarding flag set for all EC2 instances", "ports", portForwardingFlag)
				m.portForwarding = portForwardingFlag
			}
		case "ecs":
			logger.Info("Service allowed", "service", service)
			var errEcs error
			ecsClient, errEcs = GetECSClient()
			if errEcs != nil {
				logger.Fatal(errEcs)
			}
			m.focusedService = ecsService

			logger.Info("Installing the check-ecs-exec.sh script into /tmp")
			errInstallCheck := installCheck()
			if errInstallCheck != nil {
				logger.Fatal(errInstallCheck)
			}
		default:
			logger.Fatal("Service not allowed", "service", service, "allowedServices", allowedServices)
			return
		}
	}

	if len(services) == 1 {
		m.chosenService = m.focusedService
		errInitLists := m.initLists()
		if errInitLists != nil {
			m.err = errInitLists
		}
	}

	if finalModel, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		logger.Fatal(err)
	} else {
		fm := finalModel.(Model)
		if fm.command != nil {
			if fm.command.resource != nil {
				logger.Info("Command to run:")
				fmt.Println(fm.command.resource.CmdToString())
			}
		}
		if fm.err != nil {
			logger.Fatal(fm.err)
		}
	}
}
