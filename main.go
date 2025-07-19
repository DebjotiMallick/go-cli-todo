package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: todo [add|list|done|delete] [task]")
		return
	}

	command := os.Args[1]
	var taskDetail string
	if len(os.Args) > 2 {
		taskDetail = strings.Join(os.Args[2:], " ")
	}

	switch command {
	case "add":
		tasks, _ := LoadTasks()
		tasks = append(tasks, Task{Title: taskDetail, Completed: false})
		SaveTasks(tasks)
	case "list":
		ListTasks()
	case "done":
		DoneTask(taskDetail)
	case "delete":
		DeleteTask(taskDetail)
	case "clear":
		ClearTasks(taskDetail)
	case "status":
		StatusTasks()
	}
}
