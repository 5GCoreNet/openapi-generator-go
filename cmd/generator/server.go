package generator

import (
	"github.com/5GCoreNet/openapi-generator-go/pkg/generator"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Generate server code for a 5GC",
	Long:  "Generate server code for a 5GC",
	RunE: func(cmd *cobra.Command, args []string) error {
		g := &generator.Generator{
			Repository:      cmd.Flag("repository").Value.String(),
			RepositoryOwner: cmd.Flag("repository-owner").Value.String(),
			Ref:             cmd.Flag("ref").Value.String(),
			Path:            cmd.Flag("path").Value.String(),
			Output:          cmd.Flag("output").Value.String(),
		}
		return g.Generate(cmd.Context(), generator.ServerMode)
	},
}
