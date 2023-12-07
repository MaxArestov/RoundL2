package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// unpackString осуществляет примитивную распаковку строки с учетом повторений и escape-символов.
func unpackString(str string) (string, error) {
	var result string  // Строка на возврат.
	var prevRune rune  // Предыдущий символ(руна).
	isEscaped := false // bool на наличие escape-символа

	for _, val := range str {
		// Проверка, является ли символ цифрой и нет ли перед ним escape-символа.
		if unicode.IsDigit(val) && !isEscaped {
			// Возврат ошибки, если строка начинается с цифры или цифра не имеет предшествующего символа.
			if prevRune == 0 {
				return "", fmt.Errorf("некорректная строка")
			}

			count, _ := strconv.Atoi(string(val)) // Перевод считанного числа в int.
			// Добавление предыдущего символа в результирующую строку нужное количество раз.
			result += string(prevRune) + strings.Repeat(string(prevRune), count-1)
			prevRune = 0 // Сброс предыдущего символа.
		} else {
			if val == '\\' && !isEscaped { // Проверка на наличие escape-символа.
				isEscaped = true // Установка bool escape-символа.
				if prevRune != 0 {
					result += string(prevRune) // Добавление предыдущего символа в результат.
					prevRune = 0
				}
				continue
			}
			if prevRune != 0 {
				result += string(prevRune) // Добавление предыдущего символа в результат.
			}
			prevRune = val    // Установка текущего символа как предыдущего для следующей итерации.
			isEscaped = false // Сброс bool escape-символа.
		}
	}
	if prevRune != 0 {
		result += string(prevRune) // Добавление последнего символа в результирующую строку, если он есть
	}
	return result, nil // Возврат распакованной строки.
}

func main() {
	// Набор тестовых строк из ТЗ.
	testStrings := []string{"a4bc2d5e", "abcd", "45", "", `qwe\4\5`, `qwe\45`, `qwe\\5`}
	for _, val := range testStrings { // Итерация по строкам.
		unpacked, err := unpackString(val) // Распаковка строк.
		if err != nil {                    // Проверка на правильность строки.
			fmt.Printf("Принятая строка: %s\n", val)
			fmt.Printf("Ошибка: %s\n", err)
		} else {
			//Вывод принятой и распакованной строки.
			fmt.Printf("Принятая строка: %s\n", val)
			fmt.Printf("Распакованная строка: %s\n", unpacked)
		}
	}

	// Вариант с введением строки пользователем с последующей распаковкой и обработкой ошибок.
	var testString string
	fmt.Printf("Введите строку: ")

	fmt.Scan(&testString)

	unpack, err := unpackString(testString)
	if err != nil {
		fmt.Printf("Принятая строка: %s\n", testString)
		fmt.Printf("Ошибка: %s\n", err)
	} else {
		fmt.Printf("Принятая строка: %s\n", testString)
		fmt.Printf("Распакованная строка: %s\n", unpack)
	}
}
