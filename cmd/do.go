package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as completed",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			handleError(err)
			ids = append(ids, id)
		}

		tasks, err := store.GetAll()
		handleError(err)

		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number: ", id)
				continue
			}
			task := tasks[id-1]
			err = store.Delete(task.Key)
			if err != nil {
				fmt.Printf("Failed to mark \"%d\" as completed.\n", id)
				handleError(err)
			} else {
				fmt.Printf("Marked \"%d\" as completed.\n", id)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
