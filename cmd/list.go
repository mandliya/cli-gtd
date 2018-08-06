// Copyright © 2018 Ravi Mandliya <ravi.mandliya@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/GTD/db"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all your current tasks",
	Run: func(cmd *cobra.Command, args []string) {

		if db.IsTaskListEmpty() {
			fmt.Println("You have no pending tasks! Well done!")
			return
		}

		tasks, err := db.GetAllTasks()
		ExitOnError(err)

		fmt.Println("You have the following tasks:")
		for _, task := range tasks {
			fmt.Printf("%d. %s\n", task.Key, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
