package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

const dataFile = "todo.json"

func loadTasks() ([]Task, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var tasks []Task
	decErr := json.NewDecoder(file).Decode(&tasks)
	if decErr != nil {
		if decErr == io.EOF {
			return []Task{}, nil
		}
		return nil, decErr
	}
	return tasks, nil
}

func saveTasks(tasks []Task) error {
	file, err := os.Create(dataFile)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(tasks)
}

func nextID(tasks []Task) int {
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	return maxID + 1
}

func AddTask(title string) {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("加载任务失败:", err)
		return
	}
	task := Task{
		ID:    nextID(tasks),
		Title: title,
		Done:  false,
	}
	tasks = append(tasks, task)
	if err := saveTasks(tasks); err != nil {
		fmt.Println("保存任务失败:", err)
		return
	}
	fmt.Println("添加任务成功:", title)
}

func ListTasks() {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("加载任务失败:", err)
		return
	}
	if len(tasks) == 0 {
		fmt.Println("没有任务。")
		return
	}
	for _, task := range tasks {
		status := " "
		if task.Done {
			status = "✓"
		}
		fmt.Printf("%d. [%s] %s\n", task.ID, status, task.Title)
	}
}
func CompleteTask(id int) {
	panic("unimplemented")
}

func DeleteTask(id int) {
	panic("unimplemented")
}