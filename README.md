# Task Tracker CLI

## Overview
Task Tracker is a simple command-line interface (CLI) application to help you track and manage your tasks. You can add, update, delete tasks, and mark them as in-progress or done. Tasks are stored in a JSON file.

## Features
- **Add a Task**: Add a new task with a description.
- **Update a Task**: Update the description of an existing task.
- **Delete a Task**: Remove a task by its ID.
- **Mark Task as In-Progress**: Change the status of a task to "IN PROGRESS".
- **Mark Task as Done**: Change the status of a task to "DONE".
- **List Tasks**: List all tasks or filter tasks by their status.

## File Structure
task-tracker/ ├── main.go├── tasks.json├── go.mod├── utils/ │ ├── add.go│ ├── delete.go│ ├── list.go│ ├── load.go│ ├── save.go│ ├── update.go│ └── mark.go

## Usage

### Running the Application
1. **Navigate to the Project Directory**:
   ```sh
   cd path/to/task-tracker
2. go mod init task-tracker
