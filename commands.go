package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func LoadTasks() ([]Task, error) {
	var tasks []Task
	data, _ := os.ReadFile("tasks.json")
	err := json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	}
	return tasks, err
}

func ListTasks() {
	tasks, _ := LoadTasks()
	for i, task := range tasks {
		status := " "
		if task.Completed {
			status = "✅"
		} else {
			status = "❌"
		}
		fmt.Printf("%v. %v %v\n", i+1, task.Title, status)
	}
}

func SaveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
	}

	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
	return err
}

func DoneTask(taskDetail string) {
	index, err := strconv.Atoi(taskDetail)
	if err != nil {
		fmt.Println("Invalid task number:", taskDetail)
		return
	}
	index -= 1
	tasks, _ := LoadTasks()
	if index < 0 || index > len(tasks) {
		fmt.Println("Task number out of range.")
		return
	}

	if tasks[index].Completed {
		fmt.Println("Task already marked as done.")
		return
	} else {
		tasks[index].Completed = true
		fmt.Printf("Task %v marked as done", index+1)
	}

	_ = SaveTasks(tasks)
}

func DeleteTask(taskDetail string) error {
	index, err := strconv.Atoi(taskDetail)
	if err != nil {
		fmt.Println("Invalid task number:", taskDetail)
		return err
	}
	index -= 1
	tasks, _ := LoadTasks()
	if index < 0 || index >= len(tasks) {
		fmt.Println("Task number out of range.")
		return err
	}

	newTasks := append(tasks[:index], tasks[index+1:]...)
	err = SaveTasks(newTasks)
	if err == nil {
		fmt.Printf("Task %v deleted successfully", index+1)
	}
	return err
}

func ClearTasks(mode string) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	var filtered []Task

	switch mode {
	case "--done":
		for _, task := range tasks {
			if !task.Completed {
				filtered = append(filtered, task)
			}
		}
		fmt.Println("Cleared all completed tasks.")
	case "--all":
		filtered = []Task{}
		fmt.Println("Cleared all tasks.")
	}

	err = SaveTasks(filtered)
	return err
}

func StatusTasks() error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	total := len(tasks)
	completed := 0

	for _, task := range tasks {
		if task.Completed {
			completed++
		}
	}

	pending := total - completed

	fmt.Printf("You have %d tasks: %d pending, %d completed.\n", total, pending, completed)
	return nil
}
