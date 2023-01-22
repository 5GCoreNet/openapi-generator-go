package generator

import (
	"github.com/5GCoreNet/openapi-generator-go/pkg/generator"
	"github.com/spf13/cobra"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Generate client code for a 5GC",
	Long:  "Generate client code for a 5GC",
	RunE: func(cmd *cobra.Command, args []string) error {
		g := &generator.Generator{
			Repository:      cmd.Flag("repository").Value.String(),
			RepositoryOwner: cmd.Flag("repository-owner").Value.String(),
			Ref:             cmd.Flag("ref").Value.String(),
			Path:            cmd.Flag("path").Value.String(),
			Output:          cmd.Flag("output").Value.String(),
		}
		return g.Generate(cmd.Context(), generator.ClientMode)
	},
}
