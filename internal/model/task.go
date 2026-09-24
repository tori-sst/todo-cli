package model
 
import "time"

type Task struct {
	ID int `yaml:"id"`
	Title string `yaml:"title"`
	Status string `yaml:"status"`
	CreatedAt time.Time `yaml:"created_at"`
}