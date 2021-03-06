package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/dominictwlee/bao/internal/fs"
	"github.com/dominictwlee/bao/internal/logger"
	"github.com/dominictwlee/bao/internal/path"
	"github.com/dominictwlee/bao/internal/pkgjson"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
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
		/**
		1. mk project dir
		2. copy template files
		3. write pkg json according to template type
		4. install dependencies with either yarn or npm
		5. run build
		*/
		projectName := args[0]
		spin := logger.NewSpinner(logger.SpinnerOptions{
			Suffix:   logger.Title.Sprintf(" Creating project files"),
			FinalMSG: logger.Title.Sprintf("Created %v\n", projectName),
		})
		spin.Start()

		err := os.Mkdir(projectName, 0775)
		if err != nil {
			logger.Error("Failed to create project directory: %v", err)
		}

		projectPath, err := filepath.Abs(projectName)
		if err != nil {
			logger.Error("Could not find project directory path: %v", err)
		}

		modulePath, err := path.ResolveModulePath()
		if err != nil {
			logger.Error("Could not find project directory path: %v", err)
			log.Fatalln(err)
		}

		if !isValidTmpl(tmpl) {
			log.Fatalf("%s is not a valid template. Please choose from the following: %v\n", tmpl, allowedTmpls)
		}

		templatePath := filepath.Join(modulePath, "templates", strings.TrimSpace(tmpl))
		deps, ok := pkgjson.DevDepsByTmpl[tmpl]
		if !ok {
			log.Fatalf("failed to find dependencies for %s\n", tmpl)
		}

		// Copy template files
		if err := fs.CopyDir(templatePath, projectPath); err != nil {
			log.Fatalln(err)
		}

		// write pkg json
		pjson := pkgjson.New(pkgjson.NamePkg(projectName, ""))
		json, err := json.MarshalIndent(pjson, "", "\t")
		if err != nil {
			log.Fatalln(err)
		}
		if err := ioutil.WriteFile(filepath.Join(projectPath, "package.json"), json, 0664); err != nil {
			log.Fatalln(err)
		}

		spin.Stop()
		fmt.Println("")
		logger.Info("Installing Dependencies")

		// yarn/npm install dependencies
		if err := os.Chdir(projectPath); err != nil {
			log.Fatalln(err)
		}

		if err := pkgjson.InstallDeps(deps, pkgjson.DevDepFlag); err != nil {
			log.Fatalln(err)
		}

		logger.PrintProjInstructions(projectName, projectPath)
	},
}

func init() {
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	createCmd.Flags().StringVarP(&tmpl, "template", "t", "basic", "Specify a template: [basic, typescript, react, typescriptreact]. Defaults to basic")
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
