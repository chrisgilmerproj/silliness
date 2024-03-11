package main

import (
	"sort"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/charmbracelet/bubbles/list"
)

func (m *Model) initLists() {
	// Define a variable to hold the unmarshaled data
	var tagData ec2.DescribeTagsOutput
	tagData = describeTags("", "")

	// filePath := "data.json"
	// data, err := os.ReadFile(filePath)
	// if err != nil {
	// 	log.Fatal("Error reading file: %v", err)
	// }

	// // Unmarshal the JSON data into the struct
	// err = json.Unmarshal(data, &tagData)
	// if err != nil {
	// 	fmt.Printf("Error unmarshaling JSON: %v", err)
	// 	os.Exit(1)
	// }

	m.data = groupData(&tagData)

	// Init Keys
	m.lists[tagKey].Title = "Key Names"
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
	m.lists[tagKey].SetItems(keyNameItems)

	// Init Values as empty, fill this later
	m.lists[tagValue].Title = "Key Values"
	m.lists[tagValue].SetItems([]list.Item{})

	// Init Instances as empty, fill this later
	m.lists[instance].Title = "Instances"
	m.lists[instance].SetItems([]list.Item{})
}

type GroupedKeyValueData map[string]map[string][]string

func groupData(tagData *ec2.DescribeTagsOutput) GroupedKeyValueData {

	groupedData := GroupedKeyValueData{}
	for _, tagDescription := range tagData.Tags {
		// If the tag key hasn't been seen before add everything
		if _, ok := groupedData[*tagDescription.Key]; !ok {
			groupedData[*tagDescription.Key] = map[string][]string{*tagDescription.Value: {aws.ToString(tagDescription.ResourceId)}}
			continue
		}

		// If the tag key has been seen but not the value then add the list
		if _, ok := groupedData[*tagDescription.Key][*tagDescription.Value]; !ok {
			groupedData[*tagDescription.Key][*tagDescription.Value] = []string{*tagDescription.ResourceId}
			continue
		}

		// Otherwise append the resource ID
		groupedData[*tagDescription.Key][*tagDescription.Value] = append(groupedData[*tagDescription.Key][*tagDescription.Value], *tagDescription.ResourceId)
	}

	return groupedData
}
