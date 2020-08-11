package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/nahumsa/TaskManager/db"

	"github.com/mitchellh/go-homedir"
	"github.com/nahumsa/TaskManager/cmd"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	err := db.Init(dbPath)
	must(err)
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
