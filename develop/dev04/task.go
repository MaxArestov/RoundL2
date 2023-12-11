package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Функция для сортировки строки
func sortString(w string) string {
	r := []rune(w)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

// findAnagrams Функция для нахождения анаграмм
func findAnagrams(words []string) map[string][]string {
	anagrams := make(map[string][]string)
	firstWords := make(map[string]string) // Карта для хранения первых слов каждого множества

	for _, word := range words {
		lowerWord := strings.ToLower(word) // Приведение к нижнему регистру
		sortedWord := sortString(lowerWord)

		if _, exists := firstWords[sortedWord]; !exists {
			firstWords[sortedWord] = lowerWord // Сохраняем первое слово множества
		}

		anagrams[sortedWord] = append(anagrams[sortedWord], lowerWord)
	}

	// Создаем итоговую карту с правильными ключами
	finalAnagrams := make(map[string][]string)
	for _, group := range anagrams {
		if len(group) > 1 {
			firstWord := firstWords[sortString(group[0])]
			sort.Strings(group) // Сортировка группы
			finalAnagrams[firstWord] = group
		}
	}

	return finalAnagrams
}

func main() {
	words := []string{"пятак", "листок", "пятка", "слиток", "тяпка", "столик"}
	anagrams := findAnagrams(words)

	for key, group := range anagrams {
		fmt.Printf("Ключ: %s, Группа: %v\n", key, group)
	}
}
