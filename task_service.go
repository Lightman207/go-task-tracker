package main

import (
	"bufio"
	"fmt"
)

func addTask(tasks []Task, nextID int, scanner *bufio.Scanner) ([]Task, int) {
	var task Task
	fmt.Println("Введите название задачи:")
	title, err := readInput(scanner)
	if err != nil {
		fmt.Println("Введите строку")
		return tasks, nextID
	}
	task.Title = title

	fmt.Println("Введите описание задачи:")
	description, err := readInput(scanner)
	if err != nil {
		fmt.Println("Введите строку")
		return tasks, nextID
	}
	task.Description = description

	task.Status, err = readStatus(scanner)
	if err != nil {
		fmt.Println("Введите число")
		return tasks, nextID
	}

	task.ID = nextID
	nextID++
	tasks = append(tasks, task)

	return tasks, nextID
}

func showTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("Нет задач")
		return
	}
	for _, task := range tasks {
		fmt.Printf("ID: %d, Title: %s, Description: %s, Status: %s\n", task.ID, task.Title, task.Description, task.Status)
	}

}

func deleteTask(tasks []Task, scanner *bufio.Scanner) []Task {
	found := false
	fmt.Println("Введите ID задачи для удаления:")
	id, err := readInt(scanner)
	if err != nil {
		fmt.Println("Введите число")
		return tasks
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			found = true
			fmt.Println("Задача удалена")
			break
		}
	}
	if !found {
		fmt.Println("Задача с таким ID не найдена")
	}

	return tasks
}

func updateTask(tasks []Task, scanner *bufio.Scanner) []Task {
	found := false

	fmt.Println("Введите ID задачи для изменения:")
	id, err := readInt(scanner)
	if err != nil {
		fmt.Println("Введите число")
		return tasks
	}

	for i, task := range tasks {
		if task.ID == id {
			found = true
			fmt.Println("Введите новое название задачи:")
			title, err := readInput(scanner)
			if err != nil {
				fmt.Println("Введите строку")
				return tasks
			}

			fmt.Println("Введите новое описание задачи:")
			description, err := readInput(scanner)
			if err != nil {
				fmt.Println("Введите строку")
				return tasks
			}

			status, err := readStatus(scanner)
			if err != nil {
				fmt.Println("Введите число")
				return tasks
			}

			tasks[i].Status = status
			tasks[i].Title = title
			tasks[i].Description = description
			fmt.Println("Задача изменена")
			break
		}
	}
	if !found {
		fmt.Println("Задача с таким ID не найдена")
	}

	return tasks
}

func readStatus(scanner *bufio.Scanner) (string, error) {
	fmt.Println("Введите статус задачи:")
	fmt.Println("1. TODO")
	fmt.Println("2. IN PROGRESS")
	fmt.Println("3. DONE")

	status, err := readInt(scanner)
	if err != nil {
		fmt.Println("Введите число")
		return "", err
	}

	switch status {
	case 1:
		return "TODO", nil
	case 2:
		return "IN PROGRESS", nil
	case 3:
		return "DONE", nil
	default:
		return "", fmt.Errorf("неверный статус")
	}
}
