package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Вызов функции anothersolution
	anothersolution()

	in := bufio.NewScanner(os.Stdin)

	alreadySeen := make(map[string]bool)
	duplicates := []string{}

	for in.Scan() {
		txt := in.Text()

		// Проверяем на дупликат и если есть совпадение заносим значение в слайс дупликатов
		if alreadySeen[txt] {
			duplicates = append(duplicates, txt)
			continue
		}

		// Маркируем уникальное значение
		alreadySeen[txt] = true
	}

	// Вывод дупликатов: их количества и сами дупликаты
	if len(duplicates) > 0 {
		fmt.Printf("Duplicates found: %d\n", len(duplicates))
		for _, dup := range duplicates {
			fmt.Println(dup)
		}
	} else {
		fmt.Println("No duplicates found.")
	}
}

func anothersolution() {
	in := bufio.NewScanner(os.Stdin)
	var prev string

	for in.Scan() {
		txt := in.Text()
		if txt == prev {
			continue
		}
		if txt < prev {
			panic("Файл не отсортирован")
		}
		prev = txt
		fmt.Println(txt)
	}
}
