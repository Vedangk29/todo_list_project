// todo.go
package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type TodoList struct {
	Tasks []Task `json:"tasks"`
}

func (t *TodoList) AddTask(description string) {
	id := len(t.Tasks) + 1
	task := Task{ID: id, Description: description, Completed: false}
	t.Tasks = append(t.Tasks, task)
}

func (t *TodoList) CompleteTask(id int) bool {
	for i, task := range t.Tasks {
		if task.ID == id {
			t.Tasks[i].Completed = true
			return true
		}
	}
	return false
}

func (t *TodoList) SaveTasks(filename string) error {
	data, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}

func (t *TodoList) LoadTasks(filename string) error {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil
	}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, t)
}
