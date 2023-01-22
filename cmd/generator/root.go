package generator

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "generator",
		Short: "Generate code for a 5GC",
		Long:  "Generate code for a 5GC",
	}

	repository      string
	repositoryOwner string
	ref             string
	path            string
	output          string
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(clientCmd)
	rootCmd.AddCommand(serverCmd)

	rootCmd.PersistentFlags().StringVar(&repository, "repository", "", "URL to the OpenAPI specification repository")
	rootCmd.PersistentFlags().StringVar(&repositoryOwner, "repository-owner", "", "Owner of the OpenAPI specification repository")
	rootCmd.PersistentFlags().StringVar(&ref, "ref", "", "Commit hash or branch name of the OpenAPI specification repository")
	rootCmd.PersistentFlags().StringVar(&path, "path", "", "Path to the OpenAPI specification file under the repository")
	rootCmd.PersistentFlags().StringVar(&output, "output", "", "Output directory for the generated code")
}
