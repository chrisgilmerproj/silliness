package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func groupEC2Data(tagData *ec2.DescribeTagsOutput) GroupedKeyValueData {

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

func describeTags(ec2Client *ec2.Client, key, value string) ec2.DescribeTagsOutput {

	var data ec2.DescribeTagsOutput

	// Iterate through each page
	input := ec2.DescribeTagsInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("resource-type"),
				Values: []string{"instance"},
			},
		},
	}
	// Key is never used as a wildcard
	if len(key) > 0 {
		input.Filters = append(input.Filters, types.Filter{Name: aws.String("key"), Values: []string{key}})
	}
	// Value is always used as a wildcard
	if len(value) > 0 {
		value = fmt.Sprintf("*%s*", value)
		input.Filters = append(input.Filters, types.Filter{Name: aws.String("value"), Values: []string{value}})
	}
	paginator := ec2.NewDescribeTagsPaginator(ec2Client, &input)
	for paginator.HasMorePages() {
		// Get the current page
		page, err := paginator.NextPage(context.TODO())
		if err != nil {
			fmt.Println("Error getting page:", err)
		}

		data.Tags = append(data.Tags, page.Tags...)
	}
	return data
}
