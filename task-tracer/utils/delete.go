package utils

import "fmt"

// Delete a task
func DeleteTask(id int) {
    tasks := LoadTasks()
    index := -1
    for i, task := range tasks {
        if task.ID == id {
            index = i
            break
        }
    }

    if index == -1 {
        fmt.Println("Task not found.")
        return
    }

    tasks = append(tasks[:index], tasks[index+1:]...)
    SaveTasks(tasks)
    fmt.Println("Task deleted successfully!")
}
