package main

import (
	"context"

	"github.com/Azure/go-autorest/autorest/to"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	vstsrelease "github.com/microsoft/azure-devops-go-api/azuredevops/release"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	organizationUrl := "https://dev.azure.com/msazure"
	personalAccessToken := ""

	// Create a connection to your organization
	connection := azuredevops.NewPatConnection(organizationUrl, personalAccessToken)

	ctx := context.Background()

	releasClient, err := vstsrelease.NewClient(ctx, connection)
	if err != nil {
		logger.Fatalln(err)
	}

	for _, id := range []int{617, 618, 619} {
		release, err := releasClient.CreateRelease(ctx, vstsrelease.CreateReleaseArgs{
			Project: to.StringPtr("CloudNativeCompute"),
			ReleaseStartMetadata: &vstsrelease.ReleaseStartMetadata{
				DefinitionId: &id,
				Description:  to.StringPtr("Official Release 2021/06/10"),
				Artifacts: &[]vstsrelease.ArtifactMetadata{
					{
						Alias: to.StringPtr("drop_root"),
						InstanceReference: &vstsrelease.BuildVersion{
							Id:   to.StringPtr("43626541"),
							Name: to.StringPtr("v20210610.210616.3"),
						},
					},
				},
			},
		})
		if err != nil {
			logger.Warnln(err)
		} else {
			logger.Infof("%d %s %s\n", *release.Id, *release.Name)
		}
	}

	// client, err := build.NewClient(ctx, connection)
	// if err != nil {
	// 	logger.Fatalln(err)
	// }

	// m := make(map[string]string)
	// content, _ := json.Marshal(m)
	// contentStr := string(content)
	// b, err := client.QueueBuild(ctx, build.QueueBuildArgs{
	// 	Build: &build.Build{
	// 		Definition: &build.DefinitionReference{
	// 			Id: to.IntPtr(200226),
	// 		},
	// 		SourceBranch: to.StringPtr("refs/heads/official/v20210610"),
	// 		Parameters:   &contentStr,
	// 	},
	// 	Project: to.StringPtr("CloudNativeCompute"),
	// })

	// if err != nil {
	// 	logger.Fatalln(err)
	// }

	// logger.Infof("%s", b.BuildNumber)
	// logger.Infof("%s", b.BuildNumberRevision)
	// logger.Infof("%d", b.Id)
}
