package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	tasks, err := loadTasks()
	if err != nil {
		tasks = []Task{}
	}

	nextID := getNextID(tasks)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("__ Roadmap __")
		fmt.Println("1. Добавить задачу")
		fmt.Println("2. Показать задачи")
		fmt.Println("3. Изменить задачу")
		fmt.Println("4. Удалить задачу")
		fmt.Println("5. Поиск задачи")
		fmt.Println("6. Фильтр задач")
		fmt.Println("7. Выйти")

		choice, err := readInt(scanner)

		if err != nil {
			fmt.Println("Введите число")
			continue
		}

		switch choice {
		case 1:
			tasks, nextID = addTask(tasks, nextID, scanner)

			err := saveTasks(tasks)
			if err != nil {
				fmt.Println("Ошибка сохранения:", err)
			}
		case 2:
			showTasks(tasks)
		case 3:
			tasks = updateTask(tasks, scanner)

			err := saveTasks(tasks)
			if err != nil {
				fmt.Println("Ошибка сохранения:", err)
			}
		case 4:
			tasks = deleteTask(tasks, scanner)

			err := saveTasks(tasks)
			if err != nil {
				fmt.Println("Ошибка сохранения:", err)
			}
		case 5:
			searchTasks(tasks, scanner)
		case 6:
			filterTasks(tasks, scanner)
		case 7:
			fmt.Println("Выход")
			return
		default:
			fmt.Println("Некорректный выбор")
		}
	}
}
