package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ANkulagin/golang_second_brain/internal/decorator"
)

func main() {
	// Парсим аргументы командной строки
	path := flag.String("path", "", "Путь к базовой директории")
	flag.Parse()

	if *path == "" {
		log.Fatal("Пожалуйста, укажите путь к базовой директории с помощью флага -path")
	}

	// Не передаём эмодзи, оставляем пустым
	decorator := decorator.NewDecorator(*path, "")
	if err := decorator.DecorateDirectories(); err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		fmt.Println("Украшение завершено успешно.")
	}
}
