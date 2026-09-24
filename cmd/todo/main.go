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
	listCmd := flag.String("list", "", "Show task list")
	deleteCmd := flag.Int("delete", 0, "Delete task by ID")
	statusCmd := flag.Int("status", 0, "Change status by ID")
	toCmd := flag.String("to", "", "Specify new status: todo | doing | done")

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
		fmt.Printf("Task \"%s\" added successfully \n", *addCmd)
		return
	}

	if *statusCmd != 0{
		if *toCmd == ""{
			fmt.Println("Specify new status: todo | doing | done")
			return
		}

		updatedTask, err := model.UpdateStatus(tasks, *statusCmd, *toCmd)
		if err != nil{
			log.Fatalf("Error while updating")
		}
		if err := store.Save(updatedTask); err != nil{
			log.Fatalf("Error while saving: %v", err)
		}

		fmt.Printf("Task status [%d] was updated successfully \n", *statusCmd)
		return

	}

	isListSet := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == "list" {
			isListSet = true
		}
	})

	if isListSet {
		filter := *listCmd
		if filter == ""{
			filter = "all"
		}

		filteredTasks, err := model.FilteredList(tasks, filter)
		if err != nil {
			fmt.Println("Invalid filter for list")
			return
		}

		if len(filteredTasks) == 0{
			fmt.Printf("There are no tasks with status: %s \n", filter)
			return
		}

		fmt.Println("Task list: ")
		for _, t :=  range filteredTasks{
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

	fmt.Println("Tip: ")
	fmt.Println("  go run ./ -add \"Taskname\"")
	fmt.Println("  go run ./ -list [all|todo|doing|done]")
	fmt.Println("  go run ./ -status ID -to [todo|doing|done]")
	fmt.Println("  go run ./ -delete ID")
}
