package models

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

type TaskStatus string

const (
	StatusTodo       TaskStatus = "todo"
	StatusInProgress TaskStatus = "in-progress"
	StatusDone       TaskStatus = "done"
)

type Task struct {
	ID     int64      `json:"id"`
	Task   string     `json:"task"`
	Status TaskStatus `json:"status"`
}

type TaskStore struct {
	Tasks    []Task `json:"tasks"`
	mu       sync.Mutex
	filePath string
}

func NewTaskStore(filePath string) *TaskStore {
	return &TaskStore{
		filePath: filePath,
		Tasks:    []Task{},
	}
}

// GetConfigPath returns the path to the configuration file
// On Windows: %AppData%\config_task_cli\config.json
// On macOS: ~/Library/Application Support/config_task_cli/config.json
// On Linux: ~/.config/config_task_cli/config.json
// Creates the directory if it doesn't exist
func GetConfigPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	appDir := filepath.Join(configDir, "config_task_cli")
	if err := os.MkdirAll(appDir, 0755); err != nil {
		return "", err
	}
	return filepath.Join(appDir, "config.json"), nil
}

func (ts *TaskStore) Load() error {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	if _, err := os.Stat(ts.filePath); os.IsNotExist(err) {
		ts.mu.Unlock()
		err := ts.Save()
		ts.mu.Lock()
		return err
	}

	data, err := os.ReadFile(ts.filePath)
	if err != nil {
		return fmt.Errorf("error read file: %v", err)
	}

	if len(data) == 0 {
		ts.Tasks = []Task{}
		return nil
	}

	err = json.Unmarshal(data, &ts.Tasks)
	if err != nil {
		return fmt.Errorf("error parse Json: %v", err)
	}

	return nil
}

func (ts *TaskStore) Save() error {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	data, err := json.MarshalIndent(ts.Tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshal Json: %v", err)
	}

	err = os.WriteFile(ts.filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("error write file: %v", err)
	}

	return nil
}

func (ts *TaskStore) AddTask(description string) (Task, error) {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	newID := int64(1)
	if len(ts.Tasks) > 0 {
		lastTask := ts.Tasks[len(ts.Tasks)-1]
		newID = lastTask.ID + 1
	}

	task := Task{
		ID:     newID,
		Task:   description,
		Status: StatusTodo,
	}

	ts.Tasks = append(ts.Tasks, task)
	return task, nil
}

func (ts *TaskStore) UpdateStatusTask(id int64, typeStatus int64) {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	s := StatusTodo
	if typeStatus == 1 {
		s = StatusInProgress
	} else if typeStatus == 2 {
		s = StatusDone
	}

	for i := range ts.Tasks {
		if ts.Tasks[i].ID == id {
			if ts.Tasks[i].Status != s {
				ts.Tasks[i].Status = s
				break
			}
		}
	}
}

func (ts *TaskStore) UpdateTask(id int64, description string) {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	for i := range ts.Tasks {
		if ts.Tasks[i].ID == id {
			if ts.Tasks[i].Task != description {
				ts.Tasks[i].Task = description
				break
			}
		}
	}
}

func (ts *TaskStore) PrintTask(tasks []Task) {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	for _, t := range tasks {
		fmt.Println(t)
	}
}

func (ts *TaskStore) DeleteTask(id int64) []Task {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	newTasks := make([]Task, 0)

	for i := range ts.Tasks {
		if ts.Tasks[i].ID != id {
			newTasks = append(newTasks, ts.Tasks[i])
		}
	}
	return newTasks
}
