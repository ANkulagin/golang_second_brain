package decorator

import (
	"fmt"
	"github.com/ANkulagin/golang_second_brain/internal/emoji"
	"os"
	"path/filepath"
	"strings"
)

// Decorator отвечает за декорирование директорий и файлов с использованием эмодзи.
type Decorator struct {
	RootPath  string
	RootEmoji string
}

// NewDecorator создает новый экземпляр Decorator.
func NewDecorator(rootPath, rootEmoji string) *Decorator {
	return &Decorator{
		RootPath:  rootPath,
		RootEmoji: rootEmoji,
	}
}

// DecorateDirectories рекурсивно обрабатывает папки и файлы, добавляя эмодзи в имена.
func (d *Decorator) DecorateDirectories() error {
	// Если RootEmoji не задан, пробуем получить его из имени корневой директории
	if d.RootEmoji == "" {
		d.RootEmoji = emoji.GetEmoji(filepath.Base(d.RootPath))
	}

	return d.decorateDirectories(d.RootPath, d.RootEmoji)
}

func (d *Decorator) decorateDirectories(path string, inheritedEmoji string) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		oldName := entry.Name()
		oldPath := filepath.Join(path, oldName)

		// Пропускаем директории, которые начинаются с "." или "_"
		if strings.HasPrefix(oldName, ".") || strings.HasPrefix(oldName, "_") {
			continue
		}

		// Определяем эмодзи для текущего элемента: если оно уже есть, то берем его, иначе используем унаследованное.
		currentEmoji := emoji.GetEmoji(oldName)
		if currentEmoji == "" {
			currentEmoji = inheritedEmoji
		}

		// Если это директория
		if entry.IsDir() {
			// Украшаем директорию только если эмодзи отсутствует
			newName := oldName
			if inheritedEmoji != "" && emoji.GetEmoji(oldName) == "" {
				newName = emoji.AddEmoji(oldName, inheritedEmoji)
			}

			newPath := filepath.Join(path, newName)
			if oldName != newName {
				if err := os.Rename(oldPath, newPath); err != nil {
					return fmt.Errorf("failed to rename directory %s to %s: %w", oldPath, newPath, err)
				}
			}

			// Рекурсивно обрабатываем вложенные элементы с новым значением эмодзи
			if err := d.decorateDirectories(newPath, currentEmoji); err != nil {
				return err
			}
		} else { // Если это файл
			// Пропускаем файлы, не являющиеся .md
			if filepath.Ext(oldName) != ".md" {
				continue
			}

			if emoji.ContainsEmoji(oldName) {
				continue
			}

			// Разделяем имя файла и проверяем на наличие эмодзи
			fileParts := strings.SplitN(oldName, " ", 2)
			if len(fileParts) < 2 {
				// Пропускаем, если формат не соответствует ожиданиям
				continue
			}

			fileBaseName := fileParts[0]
			fileHasEmoji := emoji.GetEmoji(fileBaseName) != ""

			// Если у файла уже есть эмодзи, и оно совпадает с текущим, пропускаем
			if fileHasEmoji && emoji.GetEmoji(fileBaseName) == inheritedEmoji {
				continue
			}

			// Если у файла есть другое эмодзи, обновляем его на эмодзи родительской директории
			newName := fmt.Sprintf("%s %s %s", fileBaseName, inheritedEmoji, fileParts[1])
			newPath := filepath.Join(path, newName)

			if oldName != newName {
				if err := os.Rename(oldPath, newPath); err != nil {
					return fmt.Errorf("failed to rename file %s to %s: %w", oldPath, newPath, err)
				}
			}
		}
	}
	return nil
}
