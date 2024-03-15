package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type resourceChoice interface {
	SplitCmd() []string
	CmdToString() string
}

type ec2Choice struct {
	tag            string
	key            string
	instanceId     string
	portForwarding string
}

func (e ec2Choice) SliceCmd() []string {
	command := []string{
		"aws",
		"ssm",
		"start-session",
		"--target",
		e.instanceId,
	}
	if len(e.portForwarding) > 0 {
		portFromTo := strings.Split(e.portForwarding, ":")
		ports := map[string][]string{
			"portNumber":      {portFromTo[0]},
			"localPortNumber": {portFromTo[1]},
		}
		compactPorts, _ := json.Marshal(ports)
		command = append(command, "--document-name AWS-StartPortForwardingSession")
		command = append(command, fmt.Sprintf("--parameters '%s'", compactPorts))
	}
	return command
}

func (e ec2Choice) CmdToString() string {
	command := e.SliceCmd()
	return strings.Join(command, " ")
}

type ecsChoice struct {
	cluster       string
	containerName string
	taskId        string
}

func (e ecsChoice) SliceCmd() []string {
	command := []string{
		"aws",
		"ecs",
		"execute-command",
		"--cluster",
		e.cluster,
		"--container",
		e.containerName,
		"--task",
		e.taskId,
		"--interactive",
		"--command",
		"/bin/bash",
	}
	return command
}

func (e ecsChoice) CmdToString() string {
	command := e.SliceCmd()
	return strings.Join(command, " ")
}
