package cmd

import (
	"fmt"
	"github.com/dominictwlee/bao/internal/config"
	"github.com/dominictwlee/bao/internal/pkgjson"
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
	Long:  "Build bundles package according to your configs file and properties in package.json",
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
		_, isNode := cfg.Engines["node"]
		if pkgJson.Module != "" && !isNode {
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
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "configs", "", "configs file (default is ${ProjectRootDir}/.bao.yaml)")
}

// initConfig reads in configs file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use configs file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find working directory.
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}

		// Search configs in working directory with name ".bao" (without extension).
		viper.AddConfigPath(wd)
		viper.SetConfigName(".bao")
	}

	//viper.AutomaticEnv() // read in environment variables that match

	// If a configs file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using configs file:", viper.ConfigFileUsed())
	} else {
		fmt.Println("Could not find .bao.yml in project root. Default build options will be used")
	}
}
