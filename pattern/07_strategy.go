package pattern

/*
Можно применить:

-Когда нужно использовать разные вариации какого-то алгоритма внутри одного объекта.
-Когда есть множество похожих классов, отличающихся только некоторым поведением.
-Когда вы не хотите обнажать детали реализации алгоритмов для других классов.
-Когда различные вариации алгоритмов реализованы в виде развесистого условного оператора.
Каждая ветка такого оператора представляет собой вариацию алгоритма.

Плюсы:
-Горячая замена алгоритмов на лету.
-Изолирует код и данные алгоритмов от остальных классов.
-Уход от наследования к делегированию.
-Реализует принцип открытости/закрытости.

Минусы:
-Усложняет программу за счёт дополнительных классов.
-Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую.

Ниже представлен пример навигатора, строящего путь в зависимости от выбранного метода преодоления пути
*/

import "fmt"

type PathMode interface {
	buildPath(n *Navigator) *Path
}

type Car struct {
}

func (p *Car) buildPath(n *Navigator) *Path {
	fmt.Println("Building a path for a car")
	return &Path{
		p1: n.DeparturePoint,
		p2: n.DestinationPoint,
	}
}

type PublicTransport struct {
}

func (p *PublicTransport) buildPath(n *Navigator) *Path {
	fmt.Println("Building a path for a public transport")
	return &Path{
		p1: n.DeparturePoint,
		p2: n.DestinationPoint,
	}
}

type Foot struct {
}

func (p *Foot) buildPath(n *Navigator) *Path {
	fmt.Println("Building a path for a walk")
	return &Path{
		p1: n.DeparturePoint,
		p2: n.DestinationPoint,
	}
}

type Navigator struct {
	PathMode         PathMode
	DeparturePoint   Point
	DestinationPoint Point
}

func InitNavigator(m PathMode, departurePoint, destinationPoint Point) *Navigator {
	return &Navigator{
		PathMode:         m,
		DeparturePoint:   departurePoint,
		DestinationPoint: destinationPoint,
	}
}

func (n *Navigator) SetPathMode(m PathMode) {
	n.PathMode = m
}

func (n *Navigator) BuildPath() *Path {
	return n.PathMode.buildPath(n)
}

type Path struct {
	p1 Point
	p2 Point
}

type Point struct {
	X float32
	Y float32
}
