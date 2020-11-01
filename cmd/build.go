package cmd

import (
	"fmt"
	"github.com/dominictwlee/bao/internal/config"
	"github.com/dominictwlee/bao/internal/pkgjson"
	"github.com/evanw/esbuild/pkg/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Args:  cobra.MaximumNArgs(1),
	Short: "Build bundles package",
	Long:  "Build bundles package according to your config file and properties in package.json",
	Run: func(cmd *cobra.Command, args []string) {
		var cfg config.BuildOptions
		if err := viper.Unmarshal(&cfg); err != nil {
			log.Fatalln(err)
		}

		pkgJson, err := pkgjson.Read()
		if err != nil {
			log.Fatalln(err)
		}
		buildOpts, err := config.ConfigureBuild(&cfg, &pkgJson)
		if err != nil {
			log.Fatalln(err)
		}

		var esmOpts *api.BuildOptions
		if pkgJson.Module != "" {
			esmOpts, err = config.ConfigureESMBuild(buildOpts, &pkgJson)
		}

		allBuildOpts := []*api.BuildOptions{buildOpts, esmOpts}
		for _, opts := range allBuildOpts {
			if opts == nil {
				continue
			}
			result := api.Build(*opts)
			if len(result.Errors) > 0 {
				for _, err := range result.Errors {
					fmt.Println(err)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
