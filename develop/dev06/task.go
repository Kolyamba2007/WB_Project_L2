package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// He was extraordinarily particular\nabout6politeness6in6others

	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('/')
	if err != nil {
		panic(err)
	}

	str, _ = strings.CutSuffix(str, "/")
	fmt.Println(str)

	//result := Mancut(&line, []int{2, 4}, " ", false) //ошибка колонка меньше 1
	//fmt.Println(result)
}

func Mancut(line *string, fields []int, delimiter string, separated bool) (result []string) {
	strings1 := strings.Split(*line, "\n")
	fmt.Println(strings1)

	for _, str := range strings1 {
		if strings.Contains(str, delimiter) {
			tempStringSlice := strings.Split(str, delimiter)

			var tempString string
			for _, field := range fields {
				if len(tempStringSlice) >= field {
					tempString += tempStringSlice[field-1]
				}
			}

			result = append(result, tempString)
		} else if !separated {
			result = append(result, str)
		}
	}

	return
}
