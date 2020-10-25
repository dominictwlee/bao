package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new javascript package",
	Long:  "Create a new javascript package with all the necessary configs and scaffolding",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
	},
}

func init() {
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	createCmd.Flags().StringP("template", "t", "", "Specify a template: [typescript, react, typescriptreact]. Defaults to basic javascript")
	createCmd.AddCommand(createCmd)
}
