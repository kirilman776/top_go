package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)


func add(a, b int) int  {
    return a + b
}

func subtract(a, b int) int {
    return a - b
}

func multiply(a, b int) int {
    return a * b
}

func divide(a, b int) int {
    if b == 0 {
        fmt.Println("Ошибка: деление на ноль!")
        return 0
    }
    return a / b
}



func main() {
    var oper string
    var newtext []string
    reader := bufio.NewReader(os.Stdin)
    
    
    fmt.Println("Введите пример:")
    text, _ := reader.ReadString('\n')
    operators := [4]string{"+","-", "/", "*"}
    for _, operator := range operators {
      if strings.Contains(text, operator){
        newtext = strings.Split(strings.TrimSpace(text), operator)
        oper = operator
      }
    }
    fmt.Println(newtext)
    a, _ := strconv.Atoi(newtext[0])
    b, _ := strconv.Atoi(newtext[1])
    fmt.Println(a,b,oper)
    
    var result int

    switch oper {
    case "+":
        result = add(a, b)
    case "-":
        result = subtract(a, b)
    case "*":
        result = multiply(a, b)
    case "/":
        result = divide(a, b)
    default:
        fmt.Println("Неизвестный оператор")
        return
    }

    fmt.Printf("Результат:", result)
}
