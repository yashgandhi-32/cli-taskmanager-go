package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yashgandhi-32/testprojects/cli-taskmanager/db"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("SOmething went wrong to add task")
			return
		}
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
