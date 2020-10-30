package cmd

import (
	"github.com/evanw/esbuild/pkg/api"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Args:  cobra.MaximumNArgs(1),
	Short: "Build bundles package",
	Long:  "Build bundles package according to your config file and properties in package.json",
	Run: func(cmd *cobra.Command, args []string) {
		ioutil.WriteFile("in.ts", []byte("let x: number = 1"), 0644)

		result := api.Build(api.BuildOptions{
			Engines:     nil,
			Outfile:     "out.js",
			EntryPoints: []string{"index.js"},
			Write:       true,
			Format:      api.FormatCommonJS,
			Color:       api.ColorAlways,
		})

		if len(result.Errors) > 0 {
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
