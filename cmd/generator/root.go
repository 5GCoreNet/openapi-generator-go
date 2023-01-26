package generator

import (
	"github.com/5GCoreNet/openapi-generator-go/pkg/utils"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "api-generator-go",
		Short: "Generate code for a 5GC",
		Long:  "Generate code for a 5GC",
	}

	repository      string
	repositoryOwner string
	ref             string
	path            string
	output          string
	rootPkg         string
	exitOnFailure   bool
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(clientCmd)
	rootCmd.AddCommand(serverCmd)

	rootCmd.PersistentFlags().StringVar(&repository, "repository", utils.DefaultRepository, "URL to the OpenAPI specification repository")
	rootCmd.PersistentFlags().StringVar(&repositoryOwner, "repository-owner", utils.DefaultRepositoryOwner, "Owner of the OpenAPI specification repository")
	rootCmd.PersistentFlags().StringVar(&ref, "ref", utils.DefaultRef, "Commit hash or branch name of the OpenAPI specification repository")
	rootCmd.PersistentFlags().StringVar(&path, "path", utils.DefaultPath, "Path to the OpenAPI specification folder under the repository")
	rootCmd.PersistentFlags().StringVar(&output, "output", "", "Output directory for the generated code")
	rootCmd.PersistentFlags().StringVar(&rootPkg, "root-pkg", "", "Root package name for the generated code")
	rootCmd.PersistentFlags().BoolVar(&exitOnFailure, "exit-on-failure", true, "Exit on failure")
}
