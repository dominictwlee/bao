package logger

import (
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"os"
	"time"
)

var (
	Title = color.New(color.FgCyan, color.Bold)
)

func Error(format string, v ...interface{}) {
	color.Red(format, v...)
	os.Exit(1)
}

func Info(format string, v ...interface{}) {
	Title.Printf(format+"\n", v...)
}

type SpinnerOptions struct {
	Suffix   string
	FinalMSG string
}

func NewSpinner(opts SpinnerOptions) *spinner.Spinner {
	s := spinner.New(spinner.CharSets[7], 100*time.Millisecond)
	s.Color("fgCyan")
	if opts.Suffix != "" {
		s.Suffix = opts.Suffix
	}

	if opts.FinalMSG != "" {
		s.FinalMSG = opts.FinalMSG
	}

	return s
}

func PrintProjInstructions(projName, projPath string) {
	fmt.Printf("\nSuccessfully created %s at %s\n", projName, projPath)
	fmt.Println("Start off by navigating to your project directory with:")
	Info("cd %s\n", projName)

	fmt.Printf("Below are commands you can run in your project directory:\n\n")

	Info("yarn build")
	fmt.Printf("Build production bundles according to your project config\n\n")

	Info("yarn test")
	fmt.Printf("Test your library with Jest\n\n")
}
