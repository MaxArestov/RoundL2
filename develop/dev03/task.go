package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// flagConfig структура, содержащая все поддерживаемые флаги
type flagConfig struct {
	sortColumn   int
	numericSort  bool
	reverseSort  bool
	unique       bool
	monthSort    bool
	ignoreSpaces bool
	checkSorted  bool
	humanNumeric bool
}

/*
setupFlags устанавливает значение флагов и описывает поведение для -help, парсит
и возвращает настроенную конфигурацию.
*/
func setupFlags() *flagConfig {
	config := &flagConfig{}

	// Определяем флаги
	flag.IntVar(&config.sortColumn, "k", 0, "Указание колонки (части строки, "+
		"разделенной пробелами) для сортировки")
	flag.BoolVar(&config.numericSort, "n", false, "Сортировать по числовому значению")
	flag.BoolVar(&config.reverseSort, "r", false, "Сортировать в обратном порядке")
	flag.BoolVar(&config.unique, "u", false, "Не выводить повторяющиеся строки")
	flag.BoolVar(&config.monthSort, "M", false, "Сортировать по названию месяца")
	flag.BoolVar(&config.ignoreSpaces, "b", false, "Игнорировать хвостовые пробелы")
	flag.BoolVar(&config.checkSorted, "c", false, "Проверять, отсортированы ли данные")
	flag.BoolVar(&config.humanNumeric, "h", false, "Сортировать по числовому значению "+
		"с учетом суффиксов")

	flag.Parse()  // Парсим флаги
	return config // Возвращаем настроенный конфиг
}

// compareMonths сравнивает два значения месяца и возвращает true, если первый месяц должен идти перед вторым.
func compareMonths(valI, valJ string, reverseSort bool) bool {
	monthMap := map[string]int{
		"January": 1, "February": 2, "March": 3,
		"April": 4, "May": 5, "June": 6,
		"July": 7, "August": 8, "September": 9,
		"October": 10, "November": 11, "December": 12,
	}

	monthI, monthJ := monthMap[valI], monthMap[valJ]
	if reverseSort {
		return monthI > monthJ
	}
	return monthI < monthJ
}

// Sort принимает строки и сортирует их в соответствии с установленными флагами
func Sort(lines []string, config *flagConfig) []string {
	sort.SliceStable(lines, func(i, j int) bool {
		if config.ignoreSpaces {
			lines[i] = strings.TrimSpace(lines[i])
			lines[j] = strings.TrimSpace(lines[j])
		}

		columnsI := strings.Fields(lines[i])
		columnsJ := strings.Fields(lines[j])

		// Проверка наличия колонки
		if len(columnsI) <= config.sortColumn || len(columnsJ) <= config.sortColumn {
			return len(columnsI) < len(columnsJ)
		}

		//Присвоение значений согласно sortColumn
		valI, valJ := columnsI[config.sortColumn], columnsJ[config.sortColumn]

		if config.monthSort {
			// Используем отдельную функцию для сравнения месяцев
			return compareMonths(valI, valJ, config.reverseSort)
		}

		// Числовая сортировка, если включена
		if config.numericSort {
			var intI, intJ int
			var errI, errJ error

			// Используем функцию parseSuffixNumber, если флаг -h активен
			if config.humanNumeric {
				intI, errI = parseSuffixNumber(valI)
				intJ, errJ = parseSuffixNumber(valJ)
			} else {
				intI, errI = strconv.Atoi(valI)
				intJ, errJ = strconv.Atoi(valJ)
			}

			if errI == nil && errJ == nil {
				if config.reverseSort {
					return intI > intJ
				}
				return intI < intJ
			}
		}

		// Лексикографическая сортировка для нечисловых значений или если числовая сортировка не задана
		if config.reverseSort {
			return valI > valJ
		}
		return valI < valJ
	})

	if config.unique {
		lines = unique(lines, config.sortColumn)
	}

	return lines
}

func unique(str []string, sortColumn int) []string {
	seen := make(map[string]bool)
	result := []string{}
	for _, s := range str {
		columns := strings.Fields(s)
		var value string
		if len(columns) > sortColumn {
			value = columns[sortColumn]
		}

		if _, ok := seen[value]; !ok {
			seen[value] = true
			result = append(result, s)
		}
	}
	return result
}

// Функция для преобразования строки с числовым значением и суффиксом в число
func parseSuffixNumber(s string) (int, error) {
	multipliers := map[string]int{
		"K": 1000,
		"M": 1000000,
		"B": 1000000000,
	}

	// Проверяем, есть ли суффикс
	suffix := s[len(s)-1:]
	if multiplier, ok := multipliers[suffix]; ok {
		number, err := strconv.Atoi(s[:len(s)-1])
		if err != nil {
			return 0, err
		}
		return number * multiplier, nil
	}

	// Возвращаем обычное число, если суффикса нет
	return strconv.Atoi(s)
}

func main() {
	config := setupFlags()
	var inputFileName, outputFileName string

	fmt.Println("Введите название .txt-файла, который нужно отсортировать:")
	fmt.Scan(&inputFileName)
	outputFileName = inputFileName + "_sorted.txt"
	inputFileName += ".txt"

	// Вызываем функцию сортировки файла
	err := sortFile(inputFileName, outputFileName, config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при сортировке файла: %v\n", err)
		os.Exit(1)
	}
}

// Функция для чтения, сортировки и записи файлов
func sortFile(inputPath, outputPath string, config *flagConfig) error {
	// Открываем исходный файл
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	var lines []string

	// Читаем строки из файла
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// После чтения строк из файла проверяем на наличие ошибки сканера
	if err := scanner.Err(); err != nil {
		return err
	}

	// Сортировка строк
	sortedLines := Sort(lines, config)

	// Создаем новый файл для записи результатов
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// Записываем отсортированные строки в новый файл
	for _, line := range sortedLines {
		_, err := outputFile.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
