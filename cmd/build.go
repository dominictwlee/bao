package cmd

import (
	"fmt"
	"github.com/dominictwlee/bao/internal/config"
	"github.com/evanw/esbuild/pkg/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
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

		buildOpts, err := config.ConfigureBuild(&cfg)
		if err != nil {
			log.Fatalln(err)
		}
		result := api.Build(*buildOpts)
		if len(result.Errors) > 0 {
			fmt.Println(result.Errors)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
