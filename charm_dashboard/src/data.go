package main

import (
	"context"
	"log"
	"sort"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/charmbracelet/bubbles/list"
)

type GroupedKeyValueData map[string]map[string][]string

func (m *Model) initLists() {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// Create an Amazon S3 service client
	if m.ec2Client == nil {
		m.ec2Client = ec2.NewFromConfig(cfg)
	}
	if m.ecsClient == nil {
		m.ecsClient = ecs.NewFromConfig(cfg)
	}

	columnNames := []string{}
	switch m.chosenService {
	case ec2Service:
		var tagData ec2.DescribeTagsOutput
		tagData = describeTags(m.ec2Client, "", "")
		m.data = groupEC2Data(&tagData)
		columnNames = []string{
			"Key Names",
			"Key Values",
			"Instance IDs",
		}
	case ecsService:
		m.data = groupECSData(m.ecsClient)
		columnNames = []string{
			"Cluster Names",
			"Container Names",
			"Task IDs",
		}
	}

	// Init Keys
	m.cols[tagKey].list.Title = columnNames[0]
	var keyNameItems []list.Item

	sortedKeys := []string{}
	for key := range m.data {
		sortedKeys = append(sortedKeys, key)
	}
	sort.Strings(sortedKeys)

	for _, key := range sortedKeys {
		groupedValueData := m.data[key]
		var values []string
		for val := range groupedValueData {
			values = append(values, val)
		}
		keyNameItems = append(keyNameItems, Tag{section: tagKey, name: key, values: values})
	}
	m.cols[tagKey].list.SetItems(keyNameItems)

	// Init Values as empty, fill this later
	m.cols[tagValue].list.Title = columnNames[1]
	m.cols[tagValue].list.SetItems([]list.Item{})

	// Init resources as empty, fill this later
	m.cols[resource].list.Title = columnNames[2]
	m.cols[resource].list.SetItems([]list.Item{})

	// If there is only one available choice
	// in each category then select it and move
	if len(m.cols[tagKey].list.Items()) == 1 {
		m.SelectListItem()
	}
	if len(m.cols[tagValue].list.Items()) == 1 {
		m.SelectListItem()
	}
	if len(m.cols[resource].list.Items()) == 1 {
		m.SelectListItem()
	}
}
