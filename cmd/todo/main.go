package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/tori-sst/todo-cli/internal/model"
	"github.com/tori-sst/todo-cli/internal/storage"
)

const fileName = "tasks.yaml"

func main(){
	addCmd := flag.String("add", "", "Add a new task (name it)")
	listCmd := flag.Bool("list", false, "Show task list")
	deleteCmd := flag.Int("delete", 0, "Delete task by its ID")

	flag.Parse()

	store := storage.NewStorage(fileName)
	tasks, err := store.Load()
	if err != nil{
		log.Fatalf("Error while loading: %v", err)
	}

	if *addCmd != "" {
		tasks = model.AddTask(tasks, *addCmd)
		if err := store.Save(tasks); err != nil{
			log.Fatalf("Error while saving: %v", err)
		}
		fmt.Printf("Task \"%s\" added successful \n", *addCmd)
		return
	}

	if *listCmd {
		if len(tasks) == 0{
			fmt.Println("They're no tasks")
			return
		}
		fmt.Println("Task list: ")
		for _, t :=  range tasks{
			fmt.Printf("[%d] %s (Status: %s) - %s \n", t.ID, t.Title, t.Status, t.CreatedAt.Format("2006-01-02 15:04"))
		}
		return
	}

	if *deleteCmd != 0 {
		tasks = model.DeleteTask(tasks, *deleteCmd)
		if err := store.Save(tasks); err != nil{
			log.Fatalf("Error while saving: %v", err)
		}
		fmt.Printf("Task with ID [%d] is deleted", *deleteCmd)
		return
	}

	fmt.Println("Tip: go run ./cmd/todo -add \"Taskname\" | -list | -delete ID")
}
