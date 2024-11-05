package wiki

import (
	"os"
)

func WriteKnowledgeBase(taskLinks []string) error {
	filePath := "/home/ankul/src/obsidian/work 💼/bankiru 🏦/task 📌/wiki 📜/1 📜 Task Banki.ru Knowledge Base.md" // Укажите нужный путь
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, link := range taskLinks {
		if _, err := file.WriteString(link + "\n"); err != nil {
			return err
		}
	}
	return nil
}
