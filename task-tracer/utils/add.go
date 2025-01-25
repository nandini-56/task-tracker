package utils

import (
    "fmt"
    "time"
)

// Add a new task
func AddTask(description string) {
    tasks := LoadTasks()
    id := 1
    if len(tasks) > 0 {
        id = tasks[len(tasks)-1].ID + 1
    }
    now := time.Now()
    tasks = append(tasks, Task{
        ID:          id,
        Description: description,
        Status:      "TODO",
        CreatedAt:   now,
        UpdatedAt:   now,
    })
    SaveTasks(tasks)
    fmt.Printf("Task added successfully (ID: %d)\n", id)
}
