package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	strSlice := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {
			strSlice = append(strSlice, text)
		} else {
			break
		}
	}

	result, err := Mancut(&strSlice, []int{1, 4}, "!", true)
	if err != nil {
		panic(err)
	}

	for _, v := range result {
		fmt.Println(v)
	}
}

func Mancut(line *[]string, fields []int, delimiter string, separated bool) (result []string, err error) {
	for _, f := range fields {
		if f < 1 {
			return nil, errors.New("один из индексов меньше 1")
		}
	}

	for _, str := range *line {
		if strings.Contains(str, delimiter) {
			tempStrSlice1 := strings.Split(str, delimiter)

			var tempStrSlice2 []string
			for _, field := range fields {
				if len(tempStrSlice1) >= field {
					tempStrSlice2 = append(tempStrSlice2, tempStrSlice1[field-1])
				}
			}

			if len(tempStrSlice2) > 0 {
				result = append(result, strings.Join(tempStrSlice2, " "))
			}
		} else if !separated {
			result = append(result, str)
		}
	}

	return
}
