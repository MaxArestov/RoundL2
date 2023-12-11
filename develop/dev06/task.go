package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// cutConfig Определение структуры для конфигурации флагов
type cutConfig struct {
	fieldsFlag    string
	delimiterFlag string
	separatedFlag bool
}

// setupFlags функция для настройки и парсинга флагов
func setupFlags() *cutConfig {
	config := &cutConfig{}

	// Определение и парсинг флагов
	flag.StringVar(&config.fieldsFlag, "f", "", "выбрать поля (колонки)")
	flag.StringVar(&config.delimiterFlag, "d", "\t", "использовать другой разделитель (TAB default)")
	flag.BoolVar(&config.separatedFlag, "s", false, "только строки с разделителем")

	flag.Parse()

	return config
}

// CutLines функция для обработки строг согласно флагам
func CutLines(cfg *cutConfig, line string) string {
	// Пропуск строк без разделителя, если установлен флаг -s
	if cfg.separatedFlag && !strings.Contains(line, cfg.delimiterFlag) {
		return ""
	}

	// Разбиение строки на колонки
	columns := strings.Split(line, cfg.delimiterFlag)
	var result []string

	// Обработка и добавление указанных колонок в результат
	if cfg.fieldsFlag != "" {
		fields := strings.Split(cfg.fieldsFlag, ",")
		for _, field := range fields {
			index, err := strconv.Atoi(field)
			if err != nil || index < 1 || index > len(columns) {
				continue // Пропуск при некорректном индексе
			}
			result = append(result, columns[index-1])
		}
	} else {
		result = columns
	}
	return strings.Join(result, cfg.delimiterFlag)
}

func main() {
	config := setupFlags() // Настройка флагов и чтение конфигурации

	// Чтение и обработка строк из Stdin
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		processedLine := CutLines(config, line)
		if processedLine != "" {
			fmt.Println(processedLine)
		}
	}

	// Обработка ошибок чтения
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка чтения:", err)
	}
}
