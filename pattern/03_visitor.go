package pattern

/*
Можно применить:

-Когда нужно выполнить какую-то операцию над всеми элементами сложной структуры
объектов, например, деревом.
-Когда над объектами сложной структуры объектов надо выполнять некоторые
не связанные между собой операции, но вы не хотите «засорять» классы такими операциями.
-Когда новое поведение имеет смысл только для некоторых классов из существующей иерархии.

Плюсы:
-Упрощает добавление операций, работающих со сложными структурами объектов.
-Объединяет родственные операции в одном классе.
-Посетитель может накапливать состояние при обходе структуры элементов.

Минусы:
-Паттерн не оправдан, если иерархия элементов часто меняется.
-Может привести к нарушению инкапсуляции элементов.

Ниже представлен пример посетителя, где выполняются операции над фигурами.
В класс фигуры нужно только один раз добавить метод обработки посетителя.
Классы площади и периметра являются посетителями. Если понадобится добавить
новый функционал, то нужно всего лишь создать нового посетителя, не меняя
при этом код основного класса.
*/

import (
	"fmt"
	"math"
)

type IShape interface {
	GetShapeType() string
	Accept(IVisitor)
}

type Square struct {
	Side int
}

func (s *Square) Accept(v IVisitor) {
	v.visitForSquare(s)
}

func (s *Square) GetShapeType() string {
	return "Square"
}

type Circle struct {
	Radius int
}

func (c *Circle) Accept(v IVisitor) {
	v.visitForCircle(c)
}

func (c *Circle) GetShapeType() string {
	return "Circle"
}

type Rectangle struct {
	L int
	B int
}

func (t *Rectangle) Accept(v IVisitor) {
	v.visitForRectangle(t)
}

func (t *Rectangle) GetShapeType() string {
	return "rectangle"
}

type IVisitor interface {
	visitForSquare(*Square)
	visitForCircle(*Circle)
	visitForRectangle(*Rectangle)
}

type Area struct {
	area float32
}

func (a *Area) visitForSquare(s *Square) {
	a.area = float32(s.Side * s.Side)
	fmt.Printf("Square area: %v\n", a.area)
}

func (a *Area) visitForCircle(s *Circle) {
	a.area = math.Pi * float32(s.Radius*s.Radius)
	fmt.Printf("Circle area: %v\n", a.area)
}

func (a *Area) visitForRectangle(s *Rectangle) {
	a.area = float32(s.L * s.B)
	fmt.Printf("Rectangle area: %v\n", a.area)
}

type Perimeter struct {
	perimeter float32
}

func (a *Perimeter) visitForSquare(s *Square) {
	a.perimeter = float32(s.Side * 4)
	fmt.Printf("Square perimeter: %v\n", a.perimeter)
}

func (a *Perimeter) visitForCircle(s *Circle) {
	a.perimeter = 2 * math.Pi * float32(s.Radius)
	fmt.Printf("Circle perimeter: %v\n", a.perimeter)
}

func (a *Perimeter) visitForRectangle(s *Rectangle) {
	a.perimeter = float32(2*s.L + 2*s.B)
	fmt.Printf("Rectangle perimeter: %v\n", a.perimeter)
}
