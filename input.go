package main

import (
	"bufio"
	"fmt"
	"strconv"
)

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
