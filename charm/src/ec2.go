package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

var (
	HEALTH_STATUS_COLOR_EC2 = map[string]string{
		"pending":       "yellow",
		"running":       "green",
		"shutting-down": "magenta",
		"terminated":    "red",
	}
)

type InstanceData struct {
	InstanceId   string
	HealthStatus string
	LaunchTime   string
	Name         string
}
type GroupedValueData map[string][]InstanceData
type GroupedKeyValueData map[string]GroupedValueData

func groupEC2Instances() GroupedKeyValueData {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// Create an Amazon S3 service client
	ec2Client := ec2.NewFromConfig(cfg)

	grouped := make(GroupedKeyValueData)

	// Iterate through each page
	paginator := ec2.NewDescribeInstancesPaginator(ec2Client, &ec2.DescribeInstancesInput{})
	for paginator.HasMorePages() {
		// Get the current page
		page, err := paginator.NextPage(context.TODO())
		if err != nil {
			fmt.Println("Error getting page:", err)
		}

		// Print information about each EC2 instance on the current page
		for _, reservation := range page.Reservations {
			for _, instance := range reservation.Instances {
				if instance.State == nil {
					continue
				}
				var instanceName string
				for _, tag := range instance.Tags {
					if tag.Key == nil {
						continue
					}
					key := *tag.Key
					if tag.Value == nil {
						continue
					}
					val := *tag.Value

					// Get the name one time
					if len(instanceName) == 0 {
						for _, tag := range instance.Tags {
							if *tag.Key == "Name" {
								instanceName = val
							}
						}
					}

					if _, okKey := grouped[key]; !okKey {
						grouped[key] = GroupedValueData{val: []InstanceData{}}
					}

					instanceData := InstanceData{
						InstanceId:   aws.ToString(instance.InstanceId),
						HealthStatus: string(instance.State.Name),
						LaunchTime:   instance.LaunchTime.String(),
						Name:         instanceName,
					}
					grouped[key][val] = append(
						grouped[key][val], instanceData,
					)
				}
			}
		}
	}
	return grouped
}
