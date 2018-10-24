package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jumballaya/task/cmd"
	"github.com/jumballaya/task/db"
	homedir "github.com/mitchellh/go-homedir"
)

func main() {
	home, err := homedir.Dir()
	handleError(err)

	path := filepath.Join(home, "tasks.db")
	bucket := "tasks"

	s, err := db.New(path, bucket)
	handleError(err)
	handleError(cmd.Execute(s))
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
