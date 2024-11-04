package decorator

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ANkulagin/golang_second_brain/internal/emoji"
)

func TestDecorator_AddEmojiToDirectory(t *testing.T) {
	// Создаем временную директорию для теста
	basePath := t.TempDir()
	dirName := "TestDir"
	inheritedEmoji := "😀"

	// Создаем директорию без эмодзи
	originalPath := filepath.Join(basePath, dirName)
	if err := os.Mkdir(originalPath, 0755); err != nil {
		t.Fatalf("Failed to create directory: %v", err)
	}

	decorator := NewDecorator(basePath, inheritedEmoji)
	if err := decorator.decorateDirectories(basePath, inheritedEmoji); err != nil {
		t.Fatalf("decorateDirectories failed: %v", err)
	}

	// Проверяем, что директория была переименована с эмодзи
	expectedName := "TestDir 😀"
	expectedPath := filepath.Join(basePath, expectedName)
	if _, err := os.Stat(expectedPath); os.IsNotExist(err) {
		t.Errorf("Expected directory %s to exist, but it does not", expectedPath)
	}
}

func TestDecorator_AddEmojiToFile(t *testing.T) {
	// Создаем временную директорию для теста
	basePath := t.TempDir()
	dirName := "TestDir 😀"
	fileName := "1.1 😀 file.md"
	inheritedEmoji := ""

	// Создаем директорию с эмодзи
	originalDirPath := filepath.Join(basePath, dirName)
	if err := os.Mkdir(originalDirPath, 0755); err != nil {
		t.Fatalf("Failed to create directory: %v", err)
	}

	// Создаем файл без эмодзи
	originalFilePath := filepath.Join(originalDirPath, fileName)
	file, err := os.Create(originalFilePath)
	if err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}
	file.Close()

	decorator := NewDecorator(basePath, inheritedEmoji)
	if err := decorator.decorateDirectories(basePath, inheritedEmoji); err != nil {
		t.Fatalf("decorateDirectories failed: %v", err)
	}

	// Проверяем, что файл был переименован с эмодзи
	expectedFileName := "1.1 😀 file.md"
	expectedFilePath := filepath.Join(originalDirPath, expectedFileName)
	if _, err := os.Stat(expectedFilePath); os.IsNotExist(err) {
		t.Errorf("Expected file %s to exist, but it does not", expectedFilePath)
	}
}

func TestGetEmoji(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{"Folder 😀", "😀"},
		{"File.md", ""},
		{"Image.png 📷", "📷"},
		{"NoEmojiHere", ""},
	}

	for _, tt := range tests {
		result := emoji.GetEmoji(tt.name)
		if result != tt.expected {
			t.Errorf("GetEmoji(%q) = %q; want %q", tt.name, result, tt.expected)
		}
	}
}

func TestAddEmoji(t *testing.T) {
	tests := []struct {
		name     string
		emoji    string
		expected string
	}{
		{"Folder", "😀", "Folder 😀"},
		{"Folder 😀", "😀", "Folder 😀"},
		{"File.md", "📄", "File.md 📄"},
	}

	for _, tt := range tests {
		result := emoji.AddEmoji(tt.name, tt.emoji)
		if result != tt.expected {
			t.Errorf("AddEmoji(%q, %q) = %q; want %q", tt.name, tt.emoji, result, tt.expected)
		}
	}
}
