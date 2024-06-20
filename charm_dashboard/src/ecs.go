package main

import (
	"context"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

var (
	ecsClient *ecs.Client
	ecsOnce   sync.Once
)

// GetEC2Client returns a singleton instance of the AWS EC2 client.
func GetECSClient() (*ecs.Client, error) {
	var errDo error
	ecsOnce.Do(func() {
		var cfg aws.Config
		cfg, errDo = config.LoadDefaultConfig(context.TODO())
		ecsClient = ecs.NewFromConfig(cfg)
	})

	return ecsClient, errDo
}

func groupECSData() (GroupedKeyValueData, error) {
	data := GroupedKeyValueData{}
	ctx := context.TODO()

	listClustersPaginator := ecs.NewListClustersPaginator(ecsClient, &ecs.ListClustersInput{})
	for listClustersPaginator.HasMorePages() {
		listClustersPage, errListClusters := listClustersPaginator.NextPage(context.Background())
		if errListClusters != nil {
			return data, errListClusters
		}

		input := &ecs.DescribeClustersInput{
			Clusters: listClustersPage.ClusterArns,
		}
		describeClustersOutput, errDescribeClusters := ecsClient.DescribeClusters(ctx, input)
		if errDescribeClusters != nil {
			return data, errDescribeClusters
		}

		for _, cluster := range describeClustersOutput.Clusters {
			listTasksPaginator := ecs.NewListTasksPaginator(ecsClient, &ecs.ListTasksInput{Cluster: cluster.ClusterArn})
			for listTasksPaginator.HasMorePages() {
				listTasksPage, errListTasks := listTasksPaginator.NextPage(context.TODO())
				if errListTasks != nil {
					return data, errListTasks
				}
				describeTasksOutput, errDescribeTasks := ecsClient.DescribeTasks(ctx, &ecs.DescribeTasksInput{Cluster: cluster.ClusterArn, Tasks: listTasksPage.TaskArns})
				if errDescribeTasks != nil {
					return data, errDescribeTasks
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
							return data, errParsedArn
						}
						splitTask := strings.Split(parsedArn.Resource, "/")
						taskId := splitTask[len(splitTask)-1]
						data[clusterName][*container.Name] = append(data[clusterName][*container.Name], taskId)
					}
				}
			}
		}
	}
	return data, nil
}

func describeECSTaskHealthState(cluster string, containerName string, taskId string) (string, error) {
	ctx := context.TODO()
	describeTasksOutput, errDescribeTasks := ecsClient.DescribeTasks(ctx, &ecs.DescribeTasksInput{Cluster: &cluster, Tasks: []string{taskId}})
	if errDescribeTasks != nil {
		return "", errDescribeTasks
	}

	for _, container := range describeTasksOutput.Tasks[0].Containers {
		if *container.Name == containerName {
			return strings.ToLower(string(container.HealthStatus)), nil
		}
	}
	return "", nil
}
