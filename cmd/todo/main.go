package main

import "flag"

func main(){
	addCmd := flag.String("add", "", "Add a new task (name it)")
	listCmd := flag.Bool("list", false, "Show task list")
	deleteCmd := flag.Int("delete", 0, "Delete task by its ID")

	flag.Parse()

	if *addCmd != "" {
		println("Add task:", *addCmd)
		return
	}

	if *listCmd {
		println("Task list")
		return
	}

	if *deleteCmd != 0 {
		println("Delete task with ID: ", *deleteCmd)
		return
	}

	println("Tip: go run ./cmd/todo -add \"Taskname\" | -list | -delete ID")
}
