package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"goplay/task/db"
	"os"
)

var listCmd = &cobra.Command {
	Use: "list",
	Short: "Lists all your tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.GetAllTasks()

		if err != nil {
			fmt.Println("Something went wrong:", err)
			os.Exit(1)
		}

		if len(tasks) == 0 {
			fmt.Println("You have no tasks to complete! Why not take a vacation? 🏖️")
			return
		}
		fmt.Println("You have the following tasks:")

		for i, task := range tasks {
			fmt.Printf("%d. %s  -  %d\n", i+1, task.Value, task.Key)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}