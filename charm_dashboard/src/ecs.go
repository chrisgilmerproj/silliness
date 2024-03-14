package main

import (
	"context"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

func groupECSData() GroupedKeyValueData {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// Create an Amazon S3 service client
	ecsClient := ecs.NewFromConfig(cfg)

	data := GroupedKeyValueData{}
	ctx := context.Background()

	listClustersPaginator := ecs.NewListClustersPaginator(ecsClient, &ecs.ListClustersInput{})
	for listClustersPaginator.HasMorePages() {
		listClustersPage, errListClusters := listClustersPaginator.NextPage(context.TODO())
		if errListClusters != nil {
			log.Fatal(errListClusters)
		}

		input := &ecs.DescribeClustersInput{
			Clusters: listClustersPage.ClusterArns,
		}
		describeClustersOutput, errDescribeClusters := ecsClient.DescribeClusters(ctx, input)
		if errDescribeClusters != nil {
			log.Fatal(errDescribeClusters)
		}

		for _, cluster := range describeClustersOutput.Clusters {
			listTasksPaginator := ecs.NewListTasksPaginator(ecsClient, &ecs.ListTasksInput{Cluster: cluster.ClusterArn})
			for listTasksPaginator.HasMorePages() {
				listTasksPage, errListTasks := listTasksPaginator.NextPage(context.TODO())
				if errListTasks != nil {
					log.Fatal(errListTasks)
				}
				describeTasksOutput, errDescribeTasks := ecsClient.DescribeTasks(ctx, &ecs.DescribeTasksInput{Cluster: cluster.ClusterArn, Tasks: listTasksPage.TaskArns})
				if errDescribeTasks != nil {
					log.Fatal(errDescribeTasks)
				}

				for _, task := range describeTasksOutput.Tasks {
					for _, container := range task.Containers {

						// Check for the managed agent which is required to communicate
						managedAgentFound := false
						for _, managedAgent := range container.ManagedAgents {
							if managedAgent.Name == "ExecuteCommandAgent" {
								managedAgentFound = true
							}
						}
						if !managedAgentFound {
							continue
						}

						clusterName := *cluster.ClusterName
						if _, ok := data[clusterName]; !ok {
							data[clusterName] = map[string][]string{
								*container.Name: {},
							}
						}
						parsedArn, errParsedArn := arn.Parse(*task.TaskArn)
						if errParsedArn != nil {
							log.Fatal(errParsedArn)
						}
						splitTask := strings.Split(parsedArn.Resource, "/")
						taskId := splitTask[len(splitTask)-1]
						data[clusterName][*container.Name] = append(data[clusterName][*container.Name], taskId)
					}
				}
			}
		}
	}
	return data
}
