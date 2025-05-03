package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"time"
)

type Task struct {
	Title    string    `json:"title"`
	Message  string    `json:"message"`
	DateTime time.Time `json:"datetime"`
}

func loadTasks(filename string) ([]Task, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func notify(title, message, iconPath string) {
	cmd := exec.Command("zenity",
		"--info",
		"--title", title,
		"--text", message,
		"--ok-label", "Got it!",
		"--width", "400",
		"--height", "200",
		"--window-icon", iconPath,
	)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Failed to show zenity popup:", err)
	}
}

func main() {
	tasks, err := loadTasks("tasks.json")
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	fmt.Println("📅 Task reminder started...")

	for {
		now := time.Now()
		for i := range tasks {
			task := &tasks[i]
			if !task.DateTime.IsZero() && now.After(task.DateTime) && now.Sub(task.DateTime) < time.Second*2 {
				notify(task.Title, task.Message, "icons8-notification-500.svg")
				task.DateTime = time.Time{}
			}
		}
		time.Sleep(1 * time.Second)
	}
}
