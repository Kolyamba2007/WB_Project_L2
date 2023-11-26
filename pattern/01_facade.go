package pattern

/*
Можно применить:

-Когда вам нужно представить простой или урезанный интерфейс к сложной подсистеме.
-Когда вы хотите разложить подсистему на отдельные слои.

Плюсы:
-Изолирует клиентов от компонентов сложной подсистемы.

Минусы:
-Фасад рискует стать божественным объектом, привязанным ко всем классам программы.

Ниже представлен абстрактный пример, где мы с помощью фасада можем выполнить некоторые операции.
Но под этими операциями скрывается сложная работа с системой классов.
В реальности можно привести пример, когда мы заказываем доставку на дом.
В этом процесе участвуют десятки подсистем. В такой сложной системе легко потеряться или
что-то сломать, если обращаться с ней неправильно. Для таких случаев и существует паттерн Фасад
— он позволяет клиенту работать с десятками компонентов, используя при этом простой интерфейс.
*/

import (
	"fmt"
)

type Facade struct {
	struct1 *Struct1
	struct2 *Struct2
	struct3 *Struct3
}

func NewFacade() *Facade {
	return &Facade{
		struct1: NewStruct1(),
		struct2: NewStruct2(),
		struct3: NewStruct3(),
	}
}

func (f *Facade) Operation1() {
	fmt.Println("Operation 1")
	f.struct1.Method1()
	f.struct2.Method2()
}

func (f *Facade) Operation2() {
	fmt.Println("Operation 2")
	f.struct2.Method2()
	f.struct3.Method3()
}

type Struct1 struct {
	name string
}

func NewStruct1() *Struct1 {
	return &Struct1{name: "Struct1"}
}

func (s *Struct1) Method1() {
	fmt.Println(s.name)
}

type Struct2 struct {
	code int
}

func NewStruct2() *Struct2 {
	return &Struct2{code: 123456789}
}

func (s *Struct2) Method2() {
	fmt.Println(s.code)
}

type Struct3 struct {
	temperature float32
}

func NewStruct3() *Struct3 {
	return &Struct3{temperature: 23.6}
}

func (s *Struct3) Method3() {
	fmt.Println(s.temperature)
}
