package main

import (
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/yashgandhi-32/testprojects/cli-taskmanager/cmd"
	"github.com/yashgandhi-32/testprojects/cli-taskmanager/db"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")

	err := db.Initconn(dbPath)
	if err != nil {
		panic(err)
	}
	cmd.RootCmd.Execute()
}
