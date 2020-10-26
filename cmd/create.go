package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	tmpl string
)

var createCmd = &cobra.Command{
	Use:   "create",
	Args:  cobra.ExactArgs(1),
	Short: "Create a new javascript package",
	Long:  "Create a new javascript package with all the necessary configs and scaffolding",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		fmt.Println(cmd.Flag("template").Value)
		fmt.Println(viper.Get("hello"))
	},
}

func init() {
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	createCmd.Flags().StringVarP(&tmpl, "template", "t", "", "Specify a template: [typescript, react, typescriptreact]. Defaults to basic javascript")
	rootCmd.AddCommand(createCmd)
}
