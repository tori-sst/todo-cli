package storage

import (
	"os"

	"github.com/tori-sst/todo-cli/internal/model"
	"gopkg.in/yaml.v3"
)

type Storage struct{
	filepath string
}

func NewStorage(filepath string) *Storage{
	return &Storage{filepath: filepath}
}

func (s *Storage) Load() ([]model.Task, error) {
	if _, err := os.Stat(s.filepath); os.IsNotExist(err){
		return []model.Task{}, nil
	}

	data, err := os.ReadFile(s.filepath)
	if err != nil {
		return nil, err
	}

	var tasks []model.Task
	err = yaml.Unmarshal(data, &tasks)
	if err != nil{
		return nil, err
	}

	return tasks, nil	
}

func (s *Storage) Save(tasks []model.Task) error{
	data, err := yaml.Marshal(tasks)
	if err != nil{
		return err
	}

	return os.WriteFile(s.filepath, data, 0644)
}