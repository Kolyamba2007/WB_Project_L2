package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var strings1 []string = []string{
	"apple",
	"mangoApP",
	"watermelon",
	"cherry",
	"orangeapp",
	"cherry",
	"banana",
}

var ErrIndex = errors.New("индекс меньше 0")

func main() {
	// A2 := &After{n: 2, str: "watermelon"}
	// B2 := &Before{n: 2, str: "banana"}
	// C2 := &Context{n: 2, str: "mango"}

	// c := &Count{str: "cherry"}
	// count := GrepCount(&strings1, c)
	// fmt.Println(count)

	// i := &IgnoreCase{str: "app"}
	// v := &Invert{str: "app"}
	// F := &Fixed{str: "app"}
	n := &LineNum{str: "app"}

	grep := GrepText(&strings1, n)

	for _, str := range grep {
		fmt.Println(str)
	}
}

func GrepContext(line *[]string, key ContextKey) []string {
	result, err := key.grepContext(line)
	if err != nil {
		panic(err)
	}
	return *result
}

func GrepCount(line *[]string, key CountKey) int {
	return key.grepCount(line)
}

func GrepText(line *[]string, key TextKey) []string {
	return *key.grepText(line)
}

type ContextKey interface {
	grepContext(line *[]string) (*[]string, error)
}

type After struct {
	n   int
	str string
}

func (a *After) grepContext(line *[]string) (*[]string, error) {
	if a.n < 0 {
		return nil, ErrIndex
	}

	for i, str := range *line {
		if str == a.str {
			var N int
			if len(*line) > i+a.n {
				N = a.n + 1
			} else {
				N = len(*line) - i
			}

			result := make([]string, N)
			for j := 0; j < N; j++ {
				result[j] = (*line)[j+i]
			}

			return &result, nil
		}
	}

	return &[]string{}, nil
}

type Before struct {
	n   int
	str string
}

func (b *Before) grepContext(line *[]string) (*[]string, error) {
	if b.n < 0 {
		return nil, ErrIndex
	}

	for i, str := range *line {
		if str == b.str {
			var N int
			if b.n > i {
				N = i + 1
			} else {
				N = b.n + 1
			}

			result := make([]string, N)
			for j := 0; j < N; j++ {
				result[j] = (*line)[i-N+j+1]
			}

			return &result, nil
		}
	}

	return &[]string{}, nil
}

type Context struct {
	n   int
	str string
}

func (c *Context) grepContext(line *[]string) (*[]string, error) {
	if c.n < 0 {
		return nil, ErrIndex
	}

	for i, str := range *line {
		if str == c.str {
			var first, last int

			first = i - c.n
			if first < 0 {
				first = 0
			}

			last = i + c.n
			if last > len(*line)-1 {
				last = len(*line) - 1
			}

			N := last - first + 1

			result := make([]string, N)
			for j := 0; j < N; j++ {
				result[j] = (*line)[j+first]
			}

			return &result, nil
		}
	}

	return &[]string{}, nil
}

type CountKey interface {
	grepCount(line *[]string) int
}

type Count struct {
	str string
}

func (c *Count) grepCount(line *[]string) (result int) {
	for _, str := range *line {
		if str == c.str {
			result++
		}
	}

	return
}

type TextKey interface {
	grepText(line *[]string) *[]string
}

type IgnoreCase struct {
	str string
}

func (ic *IgnoreCase) grepText(line *[]string) *[]string {
	var result []string

	for _, str := range *line {
		var runes []rune
		for _, rune := range str {
			rune = unicode.ToLower(rune)
			runes = append(runes, rune)
		}

		if strings.Contains(string(runes), ic.str) {
			result = append(result, str)
		}
	}

	return &result
}

type Invert struct {
	str string
}

func (in *Invert) grepText(line *[]string) *[]string {
	var result []string

	for _, str := range *line {
		if !strings.Contains(str, in.str) {
			result = append(result, str)
		}
	}

	return &result
}

type Fixed struct {
	str string
}

func (f *Fixed) grepText(line *[]string) *[]string {
	var result []string

	for _, str := range *line {
		if strings.Contains(str, f.str) {
			result = append(result, str)
		}
	}

	return &result
}

type LineNum struct {
	str string
}

func (ln *LineNum) grepText(line *[]string) *[]string {
	var result []string

	for i, str := range *line {
		if strings.Contains(str, ln.str) {
			result = append(result, strconv.Itoa(i+1)+":"+str)
		}
	}

	return &result
}
