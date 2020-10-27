package cmd

import (
	"fmt"
	"github.com/dominictwlee/bao/internal/pkgjson"
	"github.com/spf13/cobra"
	"strings"
)

var (
	tmpl         string
	allowedTmpls = [...]string{"basic", "typescript", "react", "typescriptreact"}
)

var createCmd = &cobra.Command{
	Use: "create",
	Args: func(cmd *cobra.Command, args []string) error {
		err := cobra.MaximumNArgs(1)(cmd, args)
		if err != nil {
			return err
		}
		name := strings.TrimSpace(args[0])
		if _, err := pkgjson.IsValidName(name); err != nil {
			return err
		}
		return nil
	},
	Short: "Create a new javascript package",
	Long:  "Create a new javascript package with all the necessary configs and scaffolding",
	Run: func(cmd *cobra.Command, args []string) {
		if tmpl != "" {
			if !isValidTmpl(tmpl) {
				fmt.Printf("%s is not a valid template. Please choose from the following: %v\n", tmpl, allowedTmpls)
				return
			}

			// bootstrap project dir according to specified template
		} else {
			// bootstrap basic project folder
			fmt.Println("Basic")
		}
	},
}

func init() {
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	createCmd.Flags().StringVarP(&tmpl, "template", "t", "", "Specify a template: [basic, typescript, react, typescriptreact]. Defaults to basic")
	rootCmd.AddCommand(createCmd)
}

func isValidTmpl(tmpl string) bool {
	for _, t := range allowedTmpls {
		if tmpl == t {
			return true
		}
	}
	return false
}
