package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readInteger() float64 {
	var number float64
	_, err := fmt.Scanln(&number)
	if err != nil {

		fmt.Println(fmt.Sprintf("Не получилось прочитать число: %s", err))
	}
	return number
}

func readOper() string {
	reader := bufio.NewReader(os.Stdin)
	oper, _ := reader.ReadString('\n')
	oper = strings.Trim(oper, "\n")
	return oper
}

func add(a, b float64) (float64, string) {
	return a + b, "сложения"
}

func subtract(a, b float64) (float64, string) {
	return a - b, "вычитания"
}

func multiply(a, b float64) (float64, string) {
	return a * b, "умножения"
}

func divide(a, b float64) (float64, string) {
	if b == 0 {
		fmt.Println("Ошибка: деление на ноль!")
		return 0, ""
	}
	return a / b, "деления"
}

func main() {
	var result float64
	var action string

	a := readInteger()
	oper := readOper()
	b := readInteger()

	switch oper {
	case "+":
		result, action = add(a, b)
	case "-":
		result, action = subtract(a, b)
	case "*":
		result, action = multiply(a, b)
	case "/":
		result, action = divide(a, b)
		if result == 0 {
			return
		}
	default:
		fmt.Println("Неизвестный оператор")
		return
	}

	fmt.Printf("Результат %s: %.2f", action, result)
	fmt.Println()
}
