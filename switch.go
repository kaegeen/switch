package main

import (
	"fmt"
)

func main() {
	// Initialize an empty slice for the tasks
	var tasks []string

	// Display welcome message
	fmt.Println("Welcome to the To-Do List app!")
	for {
		// Display the menu
		fmt.Println("\nMenu:")
		fmt.Println("1. View tasks")
		fmt.Println("2. Add a task")
		fmt.Println("3. Remove a task")
		fmt.Println("4. Exit")
		fmt.Print("Please choose an option (1-4): ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			// View tasks
			if len(tasks) == 0 {
				fmt.Println("No tasks to display.")
			} else {
				fmt.Println("Your To-Do List:")
				for i, task := range tasks {
					fmt.Printf("%d. %s\n", i+1, task)
				}
			}

		case 2:
			// Add a task
			fmt.Print("Enter a task to add: ")
			var task string
			fmt.Scanln(&task)

			// Add task to the slice
			tasks = append(tasks, task)
			fmt.Println("Task added.")

		case 3:
			// Remove a task
			if len(tasks) == 0 {
				fmt.Println("No tasks to remove.")
				break
			}
			fmt.Print("Enter the number of the task to remove: ")
			var taskNum int
			fmt.Scanln(&taskNum)

			// Check for valid task number
			if taskNum < 1 || taskNum > len(tasks) {
				fmt.Println("Invalid task number.")
			} else {
				// Remove the task
				tasks = append(tasks[:taskNum-1], tasks[taskNum:]...)
				fmt.Println("Task removed.")
			}

		case 4:
			// Exit the program
			fmt.Println("Goodbye!")
			return

		default:
			// Handle invalid choices
			fmt.Println("Invalid option. Please choose a number between 1 and 4.")
		}
	}
}
