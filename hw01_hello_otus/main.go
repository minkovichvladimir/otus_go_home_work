package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

const SourceString = "Hello, OTUS!"

// Разумеется можно написать в одну строку.
// Добавил 2 разных действия и переменную для повышения читаемости (как условие ДЗ).
func main() {
	reversedString := reverse.String(SourceString)
	fmt.Println(reversedString)
}
