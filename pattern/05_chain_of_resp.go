package pattern

/*
Можно применить:

-Когда программа должна обрабатывать разнообразные запросы несколькими способами, но заранее
неизвестно, какие конкретно запросы будут приходить и какие обработчики для них понадобятся.
-Когда важно, чтобы обработчики выполнялись один за другим в строгом порядке.
-Когда набор объектов, способных обработать запрос, должен задаваться динамически.

Плюсы:
-Уменьшает зависимость между клиентом и обработчиками.
-Реализует принцип единственной обязанности.
-Реализует принцип открытости/закрытости.

Минусы:
-Запрос может остаться никем не обработанным.

Ниже представлен абстрактный пример, где есть запрос, посылаемый в цепочку обработчиков.
Каждый обработчик проверяет, может ли он обработать запрос, обрабатывает\необрабатывает.
Затем передает запрос следующему обработчику в цепи.

Из реальности можно привести пример, когда пользователь кликает по кнопке,
программа выстраивает цепочку из объекта этой кнопки, всех её родительских
элементов и общего окна приложения на конце. Событие клика передаётся по этой
цепи до тех пор, пока не найдётся объект, способный его обработать.
(Подход, когда обработчики прерывают цепь только когда они могут обработать запрос)
*/

import "fmt"

type Handler interface {
	Execute(*Request)
	SetNext(Handler)
}

type ConcreteHandler1 struct {
	next Handler
}

func (h *ConcreteHandler1) Execute(r *Request) {
	if r.Name == "name1" {
		r.Name = "name2"
		fmt.Println("Handler1 handled request")
		h.next.Execute(r)
		return
	}
	h.next.Execute(r)
}

func (r *ConcreteHandler1) SetNext(next Handler) {
	r.next = next
}

type ConcreteHandler2 struct {
	next Handler
}

func (h *ConcreteHandler2) Execute(r *Request) {
	if r.Value == 1 {
		r.Value = 2
		fmt.Println("Handler2 handled request")
		h.next.Execute(r)
		return
	}
	h.next.Execute(r)
}

func (r *ConcreteHandler2) SetNext(next Handler) {
	r.next = next
}

type ConcreteHandler3 struct {
	next Handler
}

func (h *ConcreteHandler3) Execute(r *Request) {
	if r.Temperature > 20.2 {
		r.Temperature *= 2
		fmt.Println("Handler3 handled request")
	}
}

func (r *ConcreteHandler3) SetNext(next Handler) {
	r.next = next
}

type Request struct {
	Name        string
	Value       int
	Temperature float32
}
