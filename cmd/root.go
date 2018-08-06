package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd is the root command
var RootCmd = &cobra.Command{
	Use:   "GTD",
	Short: "Get Things Done (GTD) is a CLI todo manager",
	Long:  "Get Things Done (GTD) is a CLI TODO manager built in Go",
}

// ExitOnError : Exits the application with an error
func ExitOnError(err error) {
	if err != nil {
		fmt.Println("Something went wrong!:", err.Error())
		os.Exit(1)
	}
}
