package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	ID          int
	Title       string
	Description string
}

func main() {
	tasks := []Task{}
	nextID := 1

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("__ Roadmap __")
		fmt.Println("1. Добавить задачу")
		fmt.Println("2. Показать задачи")
		fmt.Println("3. Удалить задачу")
		fmt.Println("4. Выйти")

		choice, err := readInt(scanner)

		if err != nil {
			fmt.Println("Введите число")
			continue
		}

		switch choice {
		case 1:
			tasks, nextID = addTask(tasks, nextID, scanner)
		case 2:
			showTasks(tasks)
		case 3:
			tasks = deleteTask(tasks, scanner)
		case 4:
			fmt.Println("Выход")
			return
		default:
			fmt.Println("Некорректный выбор")
		}
	}
}

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
		fmt.Printf("ID: %d, Title: %s, Description: %s\n", task.ID, task.Title, task.Description)
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

func readInput(scanner *bufio.Scanner) (string, error) {
	if !scanner.Scan() {
		return "", fmt.Errorf("ошибка чтения ввода")
	}
	return scanner.Text(), nil
}

func readInt(scanner *bufio.Scanner) (int, error) {
	if !scanner.Scan() {
		return 0, fmt.Errorf("ошибка чтения ввода")
	}
	return strconv.Atoi(scanner.Text())
}
