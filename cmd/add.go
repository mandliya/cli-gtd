package cmd

import (
	"fmt"
	"strings"

	"github.com/GTD/db"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		id, err := db.CreateTask(task)
		ExitOnError(err)
		fmt.Printf("Added \"%s\" to your task list with id %d \n", task, id)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
