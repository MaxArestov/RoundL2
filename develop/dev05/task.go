package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type grepConfig struct {
	printAfter      int
	printBefore     int
	printContext    int
	printLineCount  bool
	ignoreCase      bool
	invert          bool
	fixed           bool
	printLineNumber bool
}

// setupFlags устанавливает значение флагов и возвращает настроенный config
func setupFlags() *grepConfig {
	config := &grepConfig{}

	// Определяем флаги
	flag.IntVar(&config.printAfter, "A", 0, "печатать +N строк после совпадения")
	flag.IntVar(&config.printBefore, "B", 0, "печатать +N строк до совпадения")
	flag.IntVar(&config.printContext, "C", 0, "(A+B) печатать ±N строк вокруг совпадения")
	flag.BoolVar(&config.printLineCount, "c", false, "количество строк")
	flag.BoolVar(&config.ignoreCase, "i", false, "игнорировать регистр")
	flag.BoolVar(&config.invert, "v", false, "исключить вместо поиска совпадений")
	flag.BoolVar(&config.fixed, "F", false, "точное совпадение со строкой, не паттерн")
	flag.BoolVar(&config.printLineNumber, "n", false, "печатать номер строки")

	flag.Parse()

	return config
}

func GrepFile(fileName, pattern string, cfg *grepConfig) ([]string, int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, -1, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber, matchedLines := 0, 0
	var beforeLines, result []string

	for scanner.Scan() {
		line := scanner.Text()
		lineNumber++

		if cfg.ignoreCase {
			line = strings.ToLower(line)
		}

		match := strings.Contains(line, pattern)

		if cfg.fixed {
			match = line == pattern
		}

		if cfg.invert {
			match = !match
		}

		if match {
			matchedLines++
			if cfg.printBefore == 0 && cfg.printAfter == 0 && cfg.printContext == 0 {
				result = append(result, line)
			}
			if cfg.printBefore > 0 && cfg.printContext == 0 {
				result = append(result, beforeLines...)
				beforeLines = nil
				result = append(result, line)
			}
			if cfg.printAfter > 0 && cfg.printContext == 0 {
				result = append(result, line)
				for i := 0; i < cfg.printAfter; i++ {
					if scanner.Scan() {
						result = append(result, scanner.Text())
						lineNumber++
					}
				}
			}
			if cfg.printContext > 0 {
				result = append(result, beforeLines...)
				beforeLines = nil
				result = append(result, line)
				for i := 0; i < cfg.printContext; i++ {
					if scanner.Scan() {
						result = append(result, scanner.Text())
						lineNumber++
					}
				}
			}
		} else {
			if cfg.printBefore > 0 {
				if len(beforeLines) > 0 {
					if len(beforeLines) >= cfg.printBefore {
						beforeLines = beforeLines[1:]
					}
				}
			}
			if cfg.printContext > 0 {
				if len(beforeLines) > 0 {
					if len(beforeLines) >= cfg.printContext {
						beforeLines = beforeLines[1:]
					}
				}
			}
			beforeLines = append(beforeLines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, -1, err
	}

	return result, matchedLines, nil
}

func main() {
	config := setupFlags()

	if flag.NArg() < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s [flags] pattern filename\n", os.Args[0])
		os.Exit(1)
	}

	pattern := flag.Arg(0)
	fileName := flag.Arg(1)

	if config.ignoreCase {
		pattern = strings.ToLower(pattern)
	}

	lines, matchedCount, err := GrepFile(fileName, pattern, config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "grep: %v\n", err)
		os.Exit(1)
	}

	for _, line := range lines {
		fmt.Println(line)
	}
	if config.printLineCount {
		fmt.Printf(">>>>>>>Matched lines: %d<<<<<<<", matchedCount)
	}
}
