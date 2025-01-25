package main

import (
    "fmt"
    "os"
    "strconv"
    "task-tracer/utils"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: task-cli <command> [arguments]")
        return
    }

    command := os.Args[1]

    switch command {
    case "add":
        if len(os.Args) != 3 {
            fmt.Println("Usage: task-cli add <description>")
            return
        }
        description := os.Args[2]
        utils.AddTask(description)
    case "update":
        if len(os.Args) != 4 {
            fmt.Println("Usage: task-cli update <task_id> <description>")
            return
        }
        id, _ := strconv.Atoi(os.Args[2])
        description := os.Args[3]
        utils.UpdateTaskDescription(id, description)
    case "delete":
        if len(os.Args) != 3 {
            fmt.Println("Usage: task-cli delete <task_id>")
            return
        }
        id, _ := strconv.Atoi(os.Args[2])
        utils.DeleteTask(id)
    case "mark-in-progress":
        if len(os.Args) != 3 {
            fmt.Println("Usage: task-cli mark-in-progress <task_id>")
            return
        }
        id, _ := strconv.Atoi(os.Args[2])
        utils.MarkInProgress(id)
    case "mark-done":
        if len(os.Args) != 3 {
            fmt.Println("Usage: task-cli mark-done <task_id>")
            return
        }
        id, _ := strconv.Atoi(os.Args[2])
        utils.MarkDone(id)
    case "list":
        if len(os.Args) == 2 {
            utils.ListTasks("")
        } else if len(os.Args) == 3 {
            status := os.Args[2]
            utils.ListTasks(status)
        } else {
            fmt.Println("Usage: task-cli list [status]")
        }
    default:
        fmt.Println("Unknown command. Use add, update, delete, mark-in-progress, mark-done, or list.")
    }
}
