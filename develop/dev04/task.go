package main

import (
	"fmt"
	"sort"
	"unicode"
)

var str1 []string = []string{
	"тЯпка",
	"пятка",
	"Слиток",
	"листок",
	"пятак",
	"столик",
	"Тяпка",
	"слиток",
	"банан",
	"Банан",
	"банан",
}

func main() {
	result := findAnagrams(&str1)

	for key, res := range *result {
		fmt.Println(key, *res)
	}
}

func findAnagrams(words *[]string) *map[string]*[]string {
	result := make(map[string]*[]string)
	firstWords := make(map[string]string)

	for _, word := range *words {
		var runes []rune

		for _, rune := range word {
			rune = unicode.ToLower(rune)
			runes = append(runes, rune)
		}

		wordValue := string(runes)

		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})

		wordKey := string(runes)

		if _, ok := firstWords[wordKey]; !ok { // ищем строку в хэше
			firstWords[wordKey] = word // если не нашли, записываем в хэш
			(result)[word] = &[]string{wordValue}
		} else {
			*(result)[firstWords[wordKey]] = append(*(result)[firstWords[wordKey]], wordValue)
		}
	}

	for key, res := range result {
		hash := make(map[string]struct{}) // инициализация мапы
		uniqueCount := 0

		for _, str := range *res { // проходимся по строкам
			if _, ok := hash[str]; !ok { // ищем строку в хэше
				hash[str] = struct{}{} // если не нашли, записываем в хэш
				(*res)[uniqueCount] = str
				uniqueCount++
			}
		}

		for i := uniqueCount; i < len(*res); i++ {
			(*res)[i] = ""
		}

		*res = (*res)[:uniqueCount]

		if len(*res) < 2 {
			delete(result, key)
		}

		sort.Strings(*res)
	}

	return &result
}
