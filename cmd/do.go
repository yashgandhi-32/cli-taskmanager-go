package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/yashgandhi-32/testprojects/cli-taskmanager/db"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark task as complete",

	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument")
			} else {
				ids = append(ids, id)
			}
			fmt.Println(ids)
		}
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Print("Soemthing went wrong")
			return
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Print("invalid task\n")
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Failed to mark as completed error %d, %s\n", id, err)
			} else {
				fmt.Printf("Task marked as completed\n")
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
