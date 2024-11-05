package wiki

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
)

func ReadTasks(path string) ([]string, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var taskLinks []string
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".md" {
			taskLink := fmt.Sprintf("- [[%s]]", file.Name())
			taskLinks = append(taskLinks, taskLink)
		}
	}

	// Сортировка с учетом числового значения в названии файла
	sort.Slice(taskLinks, func(i, j int) bool {
		return extractFileNumber(taskLinks[i]) < extractFileNumber(taskLinks[j])
	})

	return taskLinks, nil
}

// extractFileNumber извлекает числовую часть из названия файла
func extractFileNumber(fileLink string) int {
	re := regexp.MustCompile(`\[\[(\d+)`)
	match := re.FindStringSubmatch(fileLink)
	if len(match) > 1 {
		if num, err := strconv.Atoi(match[1]); err == nil {
			return num
		}
	}
	return 0
}
