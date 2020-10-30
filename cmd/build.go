package cmd

import (
	"fmt"
	"github.com/evanw/esbuild/pkg/api"
	"github.com/spf13/cobra"
	"os"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Args:  cobra.MaximumNArgs(1),
	Short: "Build bundles package",
	Long:  "Build bundles package according to your config file and properties in package.json",
	Run: func(cmd *cobra.Command, args []string) {
		result := api.Build(api.BuildOptions{
			Color:   api.ColorAlways,
			Engines: nil,
			Outfile: "out.js",
			Loader: map[string]api.Loader{
				".js": api.LoaderJSX,
			},
			EntryPoints: []string{"sample-app/src/index.js"},
			Write:       true,
		})

		if len(result.Errors) > 0 {
			fmt.Println(result.Errors)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
