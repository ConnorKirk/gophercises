// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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
	"strings"

	"gophercise/exercise7/task/database"
	"gophercise/exercise7/task/task"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [task to add]",
	Short: "add a task to your todo list",
	Long: `add a task to your to do list.
	example: add finish task cli app`,
	Run: func(cmd *cobra.Command, args []string) {
		taskName := strings.Join(args, " ")

		//open DB
		db := database.OpenDB()
		defer db.Close()

		err := task.Add(taskName, db)
		if err != nil {
			fmt.Printf("Task could not be added: %v", err)
		}
		fmt.Printf("Task added: '%v'\n", taskName)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
