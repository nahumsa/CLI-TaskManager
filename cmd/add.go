package cmd

import (
	"fmt"
	"strings"

	"github.com/nahumsa/TaskManager/db"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task for your task list",
	Run: func(cmd *cobra.Command, args []string) {
		tsk := strings.Join(args, " ")
		_, err := db.CreateTask(tsk)
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}
		fmt.Printf("Added \"%v\" to your list\n", tsk)

	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
