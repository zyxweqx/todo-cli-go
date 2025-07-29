package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	ID        int
	Title     string
	Completed bool
}

var tasks []Task
var nextID int = 1

func main() {
	loadTask()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">>> TODO-LIST <<<")
		fmt.Println("\n1. Show tasks🔎")
		fmt.Println("2. Add new task🤝")
		fmt.Println("3. Mark the task as completed✅")
		fmt.Println("4. Delete task🗑️")
		fmt.Println("5. Save and exit❌ ")

		fmt.Print("Choose an option: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			ShowTasks()
		case "2":
			fmt.Print("Write a new task: ")
			scanner.Scan()
			title := scanner.Text()
			AddTask(title)
		case "3":
			fmt.Print("Enter ID of the task which completed: ")
			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())
			completeTask(id)
		case "4":
			fmt.Print("Enter ID of task which you wish delete: ")
			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())
			deleteTask(id)
		case "5":
			saveTasks()
			fmt.Println("Bye")
			return
		default:
			fmt.Println("No option selected❌")

		}
	}
}

func ShowTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks found❌")
		return
	}

	for _, task := range tasks {
		status := " "
		if task.Completed {
			status = "✔"
		}
		fmt.Printf("%d. [%s] %s\n", task.ID, status, task.Title)

	}
}

func AddTask(title string) {
	task := Task{
		ID:        nextID,
		Title:     title,
		Completed: false,
	}
	tasks = append(tasks, task)
	nextID++
	fmt.Println("Added new task✅")
}

func completeTask(id int) {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("Completed task✅")
			return
		}
	}
	fmt.Println("⚠️Cannot find task")
}

func deleteTask(id int) {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("🗑️Deleted task")
			return
		}
	}
	fmt.Println("⚠️Cannot find task")
}

func loadTask() {
	// Open file todo.json (os.Open)
	file, err := os.Open("todo.json")
	if err != nil {
		// Check if file is exist
		if os.IsNotExist(err) {
			return
		}
		fmt.Println("Error opening file", err)
		return
	}
	// defer
	defer file.Close()
	// decoder json
	decoder := json.NewDecoder(file)
	// read data from file and write to variable tasks decoder.Decode
	err = decoder.Decode(&tasks)
	if err != nil {
		fmt.Println("Error decoding file", err)
	}
	// Set the value of `nextID` — it must be greater than the largest task ID.
	for _, task := range tasks {
		if task.ID >= nextID {
			nextID = task.ID + 1
		}
	}
}

func saveTasks() {
	//Make or overwrite file "todo.json" (os.Create)
	file, err := os.Create("todo.json")
	if err != nil {
		fmt.Println("Error creating file", err)
		return
	}
	//defer
	defer file.Close()

	//encoder json (json.NewEncoder)
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	// write data from "tasks" to file (encoder.Encode)
	err = encoder.Encode(tasks)
	// Error or success
	if err != nil {
		fmt.Println("Error writing JSON", err)
	}
	fmt.Println("Tasks saved to", file.Name())

}
