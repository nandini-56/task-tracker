package utils

import (
    "fmt"
    "time"
)

// Update task description
func UpdateTaskDescription(id int, description string) {
    tasks := LoadTasks()
    found := false
    for i, task := range tasks {
        if task.ID == id {
            tasks[i].Description = description
            tasks[i].UpdatedAt = time.Now()
            found = true
            break
        }
    }

    if !found {
        fmt.Println("Task not found.")
        return
    }

    SaveTasks(tasks)
    fmt.Println("Task description updated successfully!")
}
