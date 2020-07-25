package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yashgandhi-32/testprojects/cli-taskmanager/db"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",

	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Print("Something went wrong")
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Print("You have no task")
			return
		}
		fmt.Print("You have following task\n")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
