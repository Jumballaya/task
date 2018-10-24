package cmd

import (
	"fmt"
	"os"

	"github.com/jumballaya/task/db"
	"github.com/spf13/cobra"
)

var store *db.Store

func handleError(err error) {
	if err != nil {
		fmt.Println("Something went wrong. ", err.Error())
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a CLI task list",
}

func Execute(s *db.Store) error {
	store = s
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
