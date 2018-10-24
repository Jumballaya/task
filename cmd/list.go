package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := store.GetAll()
		handleError(err)

		if len(tasks) > 0 {
			fmt.Println("You have the following tasks:")
			for i, t := range tasks {
				fmt.Printf("%d. %s\n", i+1, t.Value)
			}
		} else {
			fmt.Println("You are out of tasks!")
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
