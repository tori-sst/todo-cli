package model
 
import "time"

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