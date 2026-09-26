package main

import (
	"bufio"
	"fmt"
	"os"
)

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
