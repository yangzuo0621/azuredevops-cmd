package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Azure/go-autorest/autorest/to"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/git"
)

func main() {
	organizationUrl := "https://dev.azure.com/msazure"
	personalAccessToken := ""

	// Create a connection to your organization
	connection := azuredevops.NewPatConnection(organizationUrl, personalAccessToken)

	ctx := context.Background()

	// Create a client to interact with the Core area
	gitClient, err := git.NewClient(ctx, connection)
	if err != nil {
		log.Fatal(err)
	}

	result, err := gitClient.CreateAnnotatedTag(ctx, git.CreateAnnotatedTagArgs{
		RepositoryId: to.StringPtr("aks-rp"),
		Project:      to.StringPtr("CloudNativeCompute"),
		TagObject: &git.GitAnnotatedTag{
			Message: to.StringPtr("Test Tags 101..."),
			Name:    to.StringPtr("cicd.20210415.101"),
			TaggedObject: &git.GitObject{
				ObjectId: to.StringPtr("51052aa561317ef86271eae9ed84af3ea256d19d"),
			},
		},
	})

	// result, err := gitClient.GetRefs(ctx, git.GetRefsArgs{
	// 	RepositoryId:   to.StringPtr("aks-rp"),
	// 	Project:        to.StringPtr("CloudNativeCompute"),
	// 	Filter:         to.StringPtr("tags"),
	// 	FilterContains: to.StringPtr("cicd.20210414.5"),
	// })

	if err != nil {
		panic(err)
	}

	contents, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(contents))

	// _, err = gitClient.UpdateRefs(ctx, git.UpdateRefsArgs{
	// 	Project:      to.StringPtr("CloudNativeCompute"),
	// 	RepositoryId: to.StringPtr("aks-rp"),
	// 	RepositoryId: to.StringPtr("45442779-e72f-4447-9cd5-d3d2e2329e6e"),
	// 	RefUpdates: &[]git.GitRefUpdate{
	// 		{
	// 			OldObjectId: to.StringPtr("0000000000000000000000000000000000000000"),
	// 			NewObjectId: to.StringPtr("c0ad9cb0917161a26bb5e6ef710a7ee839970b38"),
	// 			Name:        to.StringPtr("refs/heads/zuya/test111"),
	// 		},
	// 	},
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// "c0ad9cb0917161a26bb5e6ef710a7ee839970b38"

	// diff, err := gitClient.GetCommitDiffs(ctx, git.GetCommitDiffsArgs{
	// 	Project:      to.StringPtr("CloudNativeCompute"),
	// 	RepositoryId: to.StringPtr("aks-rp"),
	// 	BaseVersionDescriptor: &git.GitBaseVersionDescriptor{
	// 		BaseVersion:     to.StringPtr("dab1bb614a735f5e50c9e36d6fff42cd10efef3f"),
	// 		BaseVersionType: &git.GitVersionTypeValues.Commit,
	// 	},
	// 	TargetVersionDescriptor: &git.GitTargetVersionDescriptor{
	// 		TargetVersion:     to.StringPtr("b73408a5f5b7c5ef657e9929664d565ca6fe4a30"),
	// 		TargetVersionType: &git.GitVersionTypeValues.Commit,
	// 	},
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// contents, err := json.MarshalIndent(diff, "", "  ")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Print(string(contents))

	// commits, err := gitClient.GetCommits(ctx, git.GetCommitsArgs{
	// 	Project:      to.StringPtr("CloudNativeCompute"),
	// 	RepositoryId: to.StringPtr("aks-rp"),
	// 	SearchCriteria: &git.GitQueryCommitsCriteria{
	// 		CompareVersion: &git.GitVersionDescriptor{
	// 			VersionType: &git.GitVersionTypeValues.Commit,
	// 			Version:     to.StringPtr("dab1bb614a735f5e50c9e36d6fff42cd10efef3f"),
	// 		},
	// 		ItemVersion: &git.GitVersionDescriptor{
	// 			Version:        to.StringPtr("f0c5e291b35cb06030ae400d41bf6e8ae3d83604"),
	// 			VersionType:    &git.GitVersionTypeValues.Commit,
	// 			VersionOptions: &git.GitVersionOptionsValues.FirstParent,
	// 		},
	// 	},
	// })

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, c := range *commits {
	// 	log.Printf("%s %s", *c.CommitId, *c.Comment)
	// }

	// log.Print(len(*commits))

	// commit, err := gitClient.GetCommit(ctx, git.GetCommitArgs{
	// 	CommitId:     to.StringPtr("f0c5e291b35cb06030ae400d41bf6e8ae3d83604"),
	// 	Project:      to.StringPtr("CloudNativeCompute"),
	// 	RepositoryId: to.StringPtr("aks-rp"),
	// })
	// log.Printf("%s %s", *commit.CommitId, *commit.Comment)

}
