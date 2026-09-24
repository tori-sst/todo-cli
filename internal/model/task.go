package model

import (
	"fmt"
	"time"
)

const (
	StatusTodo = "todo"
	StatusDoing = "doing"
	StatusDone = "done"
)

type Task struct {
	ID int `yaml:"id"`
	Title string `yaml:"title"`
	Status string `yaml:"status"`
	CreatedAt time.Time `yaml:"created_at"`
}

func AddTask (tasks []Task, title string) []Task{
	maxId := 0
	for _, t := range tasks {
		if t.ID > maxId{
			maxId = t.ID
		}
	}

	newTask := Task{
		ID: maxId + 1,
		Title: title,
		Status: "todo",
		CreatedAt: time.Now(),
	}

	return append(tasks, newTask)
}

func DeleteTask (tasks []Task, id int) []Task{
	var res []Task
	for _, t := range tasks {
		if t.ID != id{
			res = append(res, t)
		}
	}
	return res
}

func UpdateStatus(tasks []Task, id int, newStatus string) ([]Task, error){
	if newStatus != StatusTodo && newStatus != StatusDoing && newStatus != StatusDone{
		return nil, fmt.Errorf("Unacceptable status")
	} 

	found := false
	for t := range tasks{
		if tasks[t].ID == id {
			tasks[t].Status = newStatus
			found = true
			break
		}
	}

	if !found {
		return nil, fmt.Errorf("Task with ID [%d] is not found", id)
	}

	return tasks, nil
}

func FilteredList (tasks []Task, statusFilter string) ([]Task, error){
	if statusFilter == "all"{
		return tasks, nil
	}

	if statusFilter != StatusTodo && statusFilter != StatusDoing && statusFilter != StatusDone{
		return nil, fmt.Errorf("Unacceptable filter")
	}

	var filtered []Task
	for _, t := range tasks{
		if t.Status == statusFilter{
			filtered = append(filtered, t)
		}
	}

	return filtered, nil

}