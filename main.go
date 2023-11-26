package main

import (
	ptrn "L2/pattern"
	"fmt"
)

func main() {
	// #region Facade
	facade := ptrn.NewFacade()
	facade.Operation1()
	facade.Operation2()
	// #endregion

	// #region Builder
	cocktailBuilder := ptrn.GetBuilder("cocktail")
	recipeBuilder := ptrn.GetBuilder("recipe")

	director := ptrn.NewDirector(cocktailBuilder)
	cocktail := director.ConstructCaramelLatte()

	fmt.Printf("Caramel Latte Milk Count: %v\n", cocktail.MilkCount)
	fmt.Printf("Caramel Latte Sauce Type: %s\n", cocktail.SauceType)
	fmt.Printf("Caramel Latte Sugar: %v\n", cocktail.Sugar)

	director.SetBuilder(recipeBuilder)
	recipe := director.ConstructCaramelLatte()

	fmt.Printf("\nCaramel Latte Milk Count: %v\n", recipe.MilkCount)
	fmt.Printf("Caramel Latte Sauce Type: %s\n", recipe.SauceType)
	fmt.Printf("Caramel Latte Sugar: %v\n", recipe.Sugar)
	// #endregion

	// #region Visitor
	square := &ptrn.Square{Side: 3}
	circle := &ptrn.Circle{Radius: 4}
	rectangle := &ptrn.Rectangle{L: 4, B: 6}

	area := &ptrn.Area{}

	square.Accept(area)
	circle.Accept(area)
	rectangle.Accept(area)

	perimeter := &ptrn.Perimeter{}

	square.Accept(perimeter)
	circle.Accept(perimeter)
	rectangle.Accept(perimeter)
	// #endregion

	// #region Command
	player := &ptrn.LostfilmPlayer{}

	playCommand := &ptrn.PlayCommand{VideoPlayer: player}
	stopCommand := &ptrn.StopCommand{VideoPlayer: player}

	playButton := &ptrn.Button{Command: playCommand}
	stopButton := &ptrn.Button{Command: stopCommand}

	playButton.Handle()
	stopButton.Handle()
	// #endregion

	// #region CoR
	handler1 := &ptrn.ConcreteHandler1{}
	handler2 := &ptrn.ConcreteHandler2{}
	handler3 := &ptrn.ConcreteHandler3{}

	handler1.SetNext(handler2)
	handler2.SetNext(handler3)

	request := &ptrn.Request{Name: "name1", Value: 1, Temperature: 30.6}

	handler1.Execute(request)
	// #endregion

	// #region Factory Method
	melee, _ := ptrn.CreateEnemy("melee")
	_range, _ := ptrn.CreateEnemy("range")
	mage, _ := ptrn.CreateEnemy("mage")

	fmt.Println(melee.GetName())
	fmt.Println(melee.GetDamageValue())

	fmt.Println(_range.GetName())
	fmt.Println(_range.GetDamageValue())

	fmt.Println(mage.GetName())
	fmt.Println(mage.GetDamageValue())
	// #endregion

	// #region Strategy
	carMode := &ptrn.Car{}
	publicTransportMode := &ptrn.PublicTransport{}
	walkMode := &ptrn.Foot{}

	navigator := ptrn.Navigator{
		PathMode:         carMode,
		DeparturePoint:   ptrn.Point{X: 10.1, Y: 12.2},
		DestinationPoint: ptrn.Point{X: 60.1, Y: 102.2},
	}

	navigator.BuildPath()

	navigator.SetPathMode(publicTransportMode)
	navigator.BuildPath()

	navigator.SetPathMode(walkMode)
	path := navigator.BuildPath()

	fmt.Println(*path)
	// #endregion

	// #region State
	animator := ptrn.NewAnimator()

	animator.PlayAnim()

	animator.Transit()
	animator.PlayAnim()

	animator.SetState(animator.JumpState)
	animator.PlayAnim()
	// #endregion
}
