package utils

import (
    "encoding/json"
    "fmt"
    "os"
)

// Save tasks to JSON file
func SaveTasks(tasks []Task) {
    file, err := os.Create(TasksFile)
    if err != nil {
        fmt.Println("Error saving tasks:", err)
        return
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "  ")
    err = encoder.Encode(tasks)
    if err != nil {
        fmt.Println("Error encoding tasks:", err)
    }
}
