package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var strings1 []string = []string{
	"2 apple 3",
	"4 mango 5",
	"6 watermelon 2",
	"3 cherry 1",
	"1 orange 4",
	"5 banana 3",
	"3 cherry 1",
	"5 cherry 2",
}

func main() {
	k3 := &K{columnNumber: 3}
	r := &R{}
	u := &U{}

	Mansort(&strings1, k3, u, r)

	for _, str := range strings1 {
		fmt.Println(str)
	}
}

func Mansort(line *[]string, keys ...Key) {
	sort.Strings(*line)

	for _, key := range keys {
		key.mansort(line)
	}
}

type Key interface {
	mansort(line *[]string)
}

type K struct {
	columnNumber int
}

func (k *K) mansort(line *[]string) {
	var columns = make(map[string]string)
	var keys []string

	for i, str := range *line {
		key := strings.Split(str, " ")[k.columnNumber-1] + strconv.Itoa(i)
		columns[key] = str
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for i, key := range keys {
		(*line)[i] = columns[key]
	}
}

type N struct{}

func (*N) mansort(line *[]string) {
	var numbers []int

	for _, str := range *line {
		number, err := strconv.Atoi(str)
		if err != nil {
			break
		}

		numbers = append(numbers, number)
	}
	sort.Ints(numbers)

	strCount := len(*line) - len(numbers)
	for i := 0; i < strCount; i++ {
		(*line)[i] = (*line)[i+len(numbers)]
	}

	for i := 0; i < len(numbers); i++ {
		(*line)[i+strCount] = strconv.Itoa(numbers[i])
	}
}

type R struct{}

func (*R) mansort(line *[]string) {
	for i, j := 0, len(*line)-1; i < j; i, j = i+1, j-1 {
		(*line)[i], (*line)[j] = (*line)[j], (*line)[i] // переворачиваем строки
	}
}

type U struct{}

func (*U) mansort(line *[]string) {
	hash := make(map[string]struct{}) // инициализация мапы
	uniqueCount := 0

	for _, str := range *line { // проходимся по строкам
		if _, ok := hash[str]; !ok { // ищем строку в хэше
			hash[str] = struct{}{} // если не нашли, записываем в хэш
			(*line)[uniqueCount] = str
			uniqueCount++
		}
	}

	for i := uniqueCount; i < len(*line); i++ {
		(*line)[i] = ""
	}

	*line = (*line)[:uniqueCount]
}
