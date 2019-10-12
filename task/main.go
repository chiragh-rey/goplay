package main

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"goplay/task/cmd"
	"goplay/task/db"
	"os"
	"path/filepath"
)

// Execute: go install .
// task
// OR
// Execute: go build .
// task

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println("Something went wrong:", err)
		os.Exit(1)
	}
}