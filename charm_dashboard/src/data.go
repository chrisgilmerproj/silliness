package main

import (
	"sort"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/charmbracelet/bubbles/list"
)

type GroupedKeyValueData map[string]map[string][]string

func (m *Model) initLists() {
	columnNames := []string{}
	switch m.chosenService {
	case ec2Service:
		var tagData ec2.DescribeTagsOutput
		tagData = describeTags("", "")
		m.data = groupEC2Data(&tagData)
		columnNames = []string{
			"Key Names",
			"Key Values",
			"Instances",
		}
	case ecsService:
		m.data = groupECSData()
		columnNames = []string{
			"Cluster Names",
			"Service Names",
			"Tasks",
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

	// Init Instances as empty, fill this later
	m.cols[instance].list.Title = columnNames[2]
	m.cols[instance].list.SetItems([]list.Item{})

	// If there is only one available choice
	// in each category then select it and move
	if len(m.cols[tagKey].list.Items()) == 1 {
		m.SelectListItem()
	}
	if len(m.cols[tagValue].list.Items()) == 1 {
		m.SelectListItem()
	}
	if len(m.cols[instance].list.Items()) == 1 {
		m.SelectListItem()
	}
}
