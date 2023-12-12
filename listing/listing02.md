Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
test(): 2 - Отложенные функции могут читать и присваивать именованные возвращаемые значения возвращающей функции.
anotherTest(): 1 - Аргументы отложенной функции оцениваются при выполнении оператора defer.