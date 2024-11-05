package wiki

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadTasks(t *testing.T) {
	testDir := "./testdata"
	os.Mkdir(testDir, 0755)
	defer os.RemoveAll(testDir)

	// Создаем тестовые файлы
	testFiles := []string{"1 test1.md", "2 test2.md", "10 test10.md"}
	for _, name := range testFiles {
		file, err := os.Create(filepath.Join(testDir, name))
		if err != nil {
			t.Fatalf("Не удалось создать тестовый файл: %v", err)
		}
		file.Close()
	}

	// Ожидаемый результат
	expected := []string{"- [[1 test1.md]]", "- [[2 test2.md]]", "- [[10 test10.md]]"}

	tasks, err := ReadTasks(testDir)
	if err != nil {
		t.Fatalf("Ошибка при вызове ReadTasks: %v", err)
	}

	if len(tasks) != len(expected) {
		t.Fatalf("Ожидается %d задач, получено %d", len(expected), len(tasks))
	}

	for i, task := range tasks {
		if task != expected[i] {
			t.Errorf("Ожидается %q, получено %q", expected[i], task)
		}
	}
}
