package utils

import (
    "fmt"
    "time"
)

// Mark task as in-progress
func MarkInProgress(id int) {
    tasks := LoadTasks()
    found := false
    for i, task := range tasks {
        if task.ID == id {
            tasks[i].Status = "IN PROGRESS"
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
    fmt.Println("Task marked as in-progress!")
}

// Mark task as done
func MarkDone(id int) {
    tasks := LoadTasks()
    found := false
    for i, task := range tasks {
        if task.ID == id {
            tasks[i].Status = "DONE"
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
    fmt.Println("Task marked as done!")
}
