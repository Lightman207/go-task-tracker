package main

import "fmt"

type Task struct {
	ID          int
	Title       string
	Description string
}

func main() {
	tasks := []Task{}

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
			tasks = addTask(tasks)
		case 2:
			showTasks(tasks)
		case 3:
			fmt.Println("Удаление задачи")
		case 4:
			fmt.Println("Выход")
			return
		default:
			fmt.Println("Некорректный выбор")
		}
	}
}

func addTask(tasks []Task) []Task {
	var task Task
	fmt.Println("Введите название задачи:")
	fmt.Scan(&task.Title)

	fmt.Println("Введите описание задачи:")
	fmt.Scan(&task.Description)

	tasks = append(tasks, task)

	return tasks
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
