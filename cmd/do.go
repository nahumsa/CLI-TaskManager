package cmd

import (
	"fmt"
	"strconv"

	"github.com/nahumsa/TaskManager/db"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		// Parse args
		var ids []int
		for _, val := range args {
			id, err := strconv.Atoi(val)
			if err != nil {
				fmt.Println("Failed to parse arg:", val)
			}
			ids = append(ids, id)
		}
		fmt.Println(ids)

		// Remove
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number:", id)
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Failed to mark \"%d\" as completed. Error: %s\n", id, err)
			} else {
				fmt.Printf("Marked \"%d\" as completed.\n", id)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)

}
