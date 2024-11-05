package main

import (
	"flag"
	"fmt"
	"github.com/ANkulagin/golang_second_brain/internal/wiki"
	"log"
)

func main() {
	path := flag.String("path", "", "Путь к папке с .md файлами")
	flag.Parse()

	if *path == "" {
		log.Fatal("Пожалуйста, укажите путь с помощью флага -path.")
	}

	taskLinks, err := wiki.ReadTasks(*path)
	if err != nil {
		log.Fatalf("Ошибка при чтении задач: %v", err)
	}

	if err := wiki.WriteKnowledgeBase(taskLinks); err != nil {
		log.Fatalf("Ошибка записи в базу знаний: %v", err)
	} else {
		fmt.Println("База знаний успешно обновлена!")
	}
}
