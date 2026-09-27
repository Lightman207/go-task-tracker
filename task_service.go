package main

import (
	"bufio"
	"fmt"
	"strings"
	"time"
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

	task.Priority, err = readPriority(scanner)
	if err != nil {
		fmt.Println("Введите число")
		return tasks, nextID
	}

	task.ID = nextID
	nextID++

	now := time.Now()
	task.CreatedAt = now
	task.UpdatedAt = now

	tasks = append(tasks, task)

	return tasks, nextID
}

func showTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("Нет задач")
		return
	}
	for _, task := range tasks {
		printTask(task)
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

			priority, err := readPriority(scanner)
			if err != nil {
				fmt.Println("Введите число")
				return tasks
			}

			tasks[i].UpdatedAt = time.Now()
			tasks[i].Priority = priority
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

func readPriority(scanner *bufio.Scanner) (string, error) {
	fmt.Println("Введите приоритет задачи:")
	fmt.Println("1. LOW")
	fmt.Println("2. MEDIUM")
	fmt.Println("3. HIGH")

	priority, err := readInt(scanner)
	if err != nil {
		fmt.Println("Введите число")
		return "", err
	}

	switch priority {
	case 1:
		return "LOW", nil
	case 2:
		return "MEDIUM", nil
	case 3:
		return "HIGH", nil
	default:
		return "", fmt.Errorf("неверный приоритет")
	}
}

func searchTasks(tasks []Task, scanner *bufio.Scanner) {
	found := false
	fmt.Println("Введите название задачи для поиска:")
	title, err := readInput(scanner)
	if err != nil {
		fmt.Println("Введите строку")
		return
	}

	for _, task := range tasks {
		if strings.Contains(task.Title, title) {
			printTask(task)
			found = true
		}
	}

	if !found {
		fmt.Println("Задачи не найдены")
	}
}

func filterTasks(tasks []Task, scanner *bufio.Scanner) {
	status, err := readStatus(scanner)
	if err != nil {
		fmt.Println("Error")
		return
	}
	found := false

	for _, task := range tasks {
		if task.Status == status {
			printTask(task)
			found = true
		}
	}

	if !found {
		fmt.Println("Задачи не найдены")
	}
}

func printTask(task Task) {
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Printf("ID: %d\n", task.ID)
	fmt.Printf("Название: %s\n", task.Title)
	fmt.Printf("Описание: %s\n", task.Description)
	fmt.Printf("Статус: %s\n", task.Status)
	fmt.Printf("Приоритет: %s\n", task.Priority)
	fmt.Printf("Создана: %s\n", task.CreatedAt.Format("02.01.2006 15:04"))
	fmt.Printf("Изменена: %s\n", task.UpdatedAt.Format("02.01.2006 15:04"))
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
}
