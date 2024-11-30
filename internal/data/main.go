package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Структура для хранения данных о расходах по категориям
type Expense struct {
	Category string
	Andrey   int
	Yulia    int
}

func main() {
	// Путь к папке с файлами
	dir := "/home/ankul/obsidian/_notes/daily"

	// Инициализируем переменные для хранения общих сумм и данных по категориям
	totalAndrey := 0
	totalYulia := 0
	categoryExpenses := make(map[string][2]int) // Карта для хранения расходов по категориям для Андрея и Юли

	// Проходим по всем файлам в указанной директории
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Проверяем, что это файл и его имя оканчивается на ".md"
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			file, err := os.Open(path)
			if err != nil {
				fmt.Println("Ошибка открытия файла:", err)
				return nil
			}
			defer file.Close()

			// Читаем файл построчно
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := scanner.Text()

				// Проверяем, является ли строка данными таблицы
				if strings.HasPrefix(line, "|") {
					columns := strings.Split(line, "|")
					if len(columns) >= 3 {
						category := strings.TrimSpace(columns[1])
						andreyAmount, _ := strconv.Atoi(strings.TrimSpace(columns[2]))
						yuliaAmount, _ := strconv.Atoi(strings.TrimSpace(columns[3]))

						// Игнорируем строки заголовков и разделителей
						if category == "Category" || strings.Contains(category, "-") {
							continue
						}

						// Суммируем расходы по категориям
						categoryExpenses[category] = [2]int{
							categoryExpenses[category][0] + andreyAmount,
							categoryExpenses[category][1] + yuliaAmount,
						}

						// Суммируем общие расходы
						totalAndrey += andreyAmount
						totalYulia += yuliaAmount
					}
				}
			}
			if err := scanner.Err(); err != nil {
				fmt.Println("Ошибка чтения файла:", err)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println("Ошибка обхода директории:", err)
		return
	}

	// Выводим результаты для каждого участника
	fmt.Printf("Андрей потратил: %d -> ", totalAndrey)
	for category, amounts := range categoryExpenses {
		if amounts[0] > 0 {
			fmt.Printf("%s: %d, ", category, amounts[0])
		}
	}
	fmt.Println()

	fmt.Printf("Юля потратила: %d -> ", totalYulia)
	for category, amounts := range categoryExpenses {
		if amounts[1] > 0 {
			fmt.Printf("%s: %d, ", category, amounts[1])
		}
	}
	fmt.Println()

	// Выводим общую сумму
	fmt.Printf("\nВсего потрачено: %d\n", totalAndrey+totalYulia)
}
