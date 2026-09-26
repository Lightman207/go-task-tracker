package main

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	ID          int
	Title       string
	Description string
}

func main() {
	tasks := []Task{}
	nextID := 1

	for {
		fmt.Println("__ Roadmap __")
		fmt.Println("1. Добавить задачу")
		fmt.Println("2. Показать задачи")
		fmt.Println("3. Удалить задачу")
		fmt.Println("4. Выйти")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			tasks, nextID = addTask(tasks, nextID)
		case 2:
			showTasks(tasks)
		case 3:
			tasks = deleteTask(tasks)
		case 4:
			fmt.Println("Выход")
			return
		default:
			fmt.Println("Некорректный выбор")
		}
	}
}

func addTask(tasks []Task, nextID int) ([]Task, int) {
	var task Task

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите название задачи:")
	scanner.Scan()
	task.Title = scanner.Text()

	fmt.Println("Введите описание задачи:")
	scanner.Scan()
	task.Description = scanner.Text()

	task.ID = nextID
	nextID++
	tasks = append(tasks, task)

	return tasks, nextID
}

func showTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("Нет задач")
		return
	} else {
		for _, task := range tasks {
			fmt.Printf("ID: %d, Title: %s, Description: %s\n", task.ID, task.Title, task.Description)
		}
	}
}

func deleteTask(tasks []Task) []Task {
	var id int
	found := false
	fmt.Println("Введите ID задачи для удаления:")
	fmt.Scan(&id)

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
