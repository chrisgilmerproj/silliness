package main

import (
	"sort"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/charmbracelet/bubbles/list"
)

type GroupedKeyValueData map[string]map[string][]string

func (m *Model) initLists() error {

	columnNames := []string{}
	switch m.chosenService {
	case ec2Service:
		var tagData ec2.DescribeTagsOutput
		var errDescribeTags error
		tagData, errDescribeTags = describeTags("", "")
		if errDescribeTags != nil {
			return errDescribeTags
		}
		m.data = groupEC2Data(&tagData)
		columnNames = []string{
			"Key Names",
			"Key Values",
			"Instance IDs",
		}
	case ecsService:
		var errGroupECSData error
		m.data, errGroupECSData = groupECSData()
		if errGroupECSData != nil {
			return errGroupECSData
		}
		columnNames = []string{
			"Cluster Names",
			"Container Names",
			"Task IDs",
		}
	}

	// Init Keys
	m.columns[tagKey].list.Title = columnNames[0]
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
		keyNameItems = append(keyNameItems, SectionItem{section: tagKey, name: key, values: values})
	}
	m.columns[tagKey].list.SetItems(keyNameItems)

	// Init Values as empty, fill this later
	m.columns[tagValue].list.Title = columnNames[1]
	m.columns[tagValue].list.SetItems([]list.Item{})

	// Init resources as empty, fill this later
	m.columns[resource].list.Title = columnNames[2]
	m.columns[resource].list.SetItems([]list.Item{})

	// If there is only one available choice
	// in each category then select it and move
	if len(m.columns[tagKey].list.Items()) == 1 {
		m.SelectListItem()
	}
	if len(m.columns[tagValue].list.Items()) == 1 {
		m.SelectListItem()
	}
	if len(m.columns[resource].list.Items()) == 1 {
		m.SelectListItem()
	}

	return nil
}
