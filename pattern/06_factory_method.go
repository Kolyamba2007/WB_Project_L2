package pattern

/*
Можно применить:

-Когда заранее неизвестны типы и зависимости объектов, с которыми должен работать код.
-Когда вы хотите дать возможность пользователям расширять части вашего фреймворка или библиотеки.
-Когда вы хотите экономить системные ресурсы, повторно используя уже созданные объекты, вместо порождения новых.

Плюсы:
-Избавляет класс от привязки к конкретным классам продуктов.
-Выделяет код производства продуктов в одно место, упрощая поддержку кода.
-Упрощает добавление новых продуктов в программу.
-Реализует принцип открытости/закрытости.

Минусы:
-Может привести к созданию больших параллельных иерархий классов,
так как для каждого класса продукта надо создать свой подкласс создателя.

Ниже представлен пример фабрики, которая порождает врагов с именем и уроном в зависимости от их типа
*/

import (
	"errors"
)

type IEnemy interface {
	SetName(name string)
	SetDamageValue(power int)
	GetName() string
	GetDamageValue() int
}

type Enemy struct {
	name        string
	damageValue int
}

func (g *Enemy) SetName(name string) {
	g.name = name
}

func (g *Enemy) GetName() string {
	return g.name
}

func (g *Enemy) SetDamageValue(damageValue int) {
	g.damageValue = damageValue
}

func (g *Enemy) GetDamageValue() int {
	return g.damageValue
}

type Melee struct {
	Enemy
}

func newMelee() IEnemy {
	return &Melee{
		Enemy: Enemy{
			name:        "Melee",
			damageValue: 10,
		},
	}
}

type Range struct {
	Enemy
}

func newRange() IEnemy {
	return &Range{
		Enemy: Enemy{
			name:        "Range",
			damageValue: 6,
		},
	}
}

type Mage struct {
	Enemy
}

func newMage() IEnemy {
	return &Mage{
		Enemy: Enemy{
			name:        "Mage",
			damageValue: 14,
		},
	}
}

func CreateEnemy(enemyType string) (IEnemy, error) {
	switch enemyType {
	case "melee":
		return newMelee(), nil
	case "range":
		return newRange(), nil
	case "mage":
		return newMage(), nil
	default:
		return nil, errors.New("wrong enemy type passed")
	}
}
