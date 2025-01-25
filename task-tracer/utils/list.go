package utils

import "fmt"

func ListTasks(filter string) {
    tasks := LoadTasks()
    var filteredTasks []Task

    for _, task := range tasks {
        if filter == "" || task.Status == filter {
            filteredTasks = append(filteredTasks, task)
        }
    }

    if len(filteredTasks) == 0 {
        fmt.Println("No tasks found.")
        return
    }

    for _, task := range filteredTasks {
        fmt.Printf("ID: %d\nDescription: %s\nStatus: %s\nCreated At: %s\nUpdated At: %s\n\n",
            task.ID, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
    }
}