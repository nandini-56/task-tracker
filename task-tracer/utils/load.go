package utils

import (
    "encoding/json"
    "fmt"
    "io"
    "os"
    "time"
)

const TasksFile = "tasks.json"

type Task struct {
    ID          int       `json:"id"`
    Description string    `json:"description"`
    Status      string    `json:"status"`
    CreatedAt   time.Time `json:"createdAt"`
    UpdatedAt   time.Time `json:"updatedAt"`
}

// Load tasks from JSON file
func LoadTasks() []Task {
    if _, err := os.Stat(TasksFile); os.IsNotExist(err) {
        os.WriteFile(TasksFile, []byte("[]"), 0644)
    }

    file, err := os.Open(TasksFile)
    if err != nil {
        fmt.Println("Error reading tasks file:", err)
        return []Task{}
    }
    defer file.Close()

    var tasks []Task
    decoder := json.NewDecoder(file)
    err = decoder.Decode(&tasks)
    if err != nil && err != io.EOF {
        fmt.Println("Error decoding tasks file:", err)
    }

    return tasks
}
