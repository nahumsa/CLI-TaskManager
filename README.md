# CLI Task Manager

This is a Command Line task manager based on [gophercises](https://gophercises.com/) which uses [BoltDB](https://github.com/boltdb/bolt) to store tasks and uses [Cobra](https://github.com/spf13/cobra) to generate the commands. 

To run in your personal computer you only need use `go install .`, then just need to call `TaskManager` on your command line and you're ready to go!

The available commands are:
  - add:         Adds a task for your task list
  - do:          Marks a task as complete
  - help:        Help about any command
  - list:        Lists all your tasks