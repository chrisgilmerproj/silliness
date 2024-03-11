package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go/aws"
)

func describeTags(key, value string) ec2.DescribeTagsOutput {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// Create an Amazon S3 service client
	ec2Client := ec2.NewFromConfig(cfg)

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
