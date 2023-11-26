package pattern

/*
Можно применить:

-Когда есть объект, поведение которого кардинально меняется в зависимости
от внутреннего состояния, причём типов состояний много, и их код часто меняется.
-Когда код класса содержит множество больших, похожих друг на друга, условных
операторов, которые выбирают поведения в зависимости от текущих значений полей класса.

Плюсы:
-Избавляет от множества больших условных операторов машины состояний.
-Концентрирует в одном месте код, связанный с определённым состоянием.
-Упрощает код контекста.

Минусы:
-Может неоправданно усложнить код, если состояний мало и они редко меняются.

Ниже представлен пример аниматора, проигрывающего анимации в зависимости от текущего состояния
*/

import "fmt"

type Animator struct {
	RunState  State
	JumpState State
	WalkState State

	currentState State
}

func NewAnimator() *Animator {
	a := &Animator{}

	a.RunState = &Run{animator: a}
	a.JumpState = &Jump{animator: a}
	a.WalkState = &Walk{animator: a}
	a.SetState(a.WalkState)

	return a
}

func (a *Animator) SetState(s State) {
	a.currentState = s
}

func (a *Animator) PlayAnim() {
	a.currentState.playAnim()
}

func (a *Animator) Transit() {
	a.currentState.transit()
}

type State interface {
	playAnim()
	transit()
}

type Run struct {
	animator *Animator
}

func (r *Run) playAnim() {
	fmt.Println("Playing run animation")
}

func (r *Run) transit() {
	fmt.Println("Transition to walking state")
	r.animator.SetState(r.animator.WalkState)
}

type Jump struct {
	animator *Animator
}

func (j *Jump) playAnim() {
	fmt.Println("Playing jump animation")
}

func (j *Jump) transit() {
	fmt.Println("Transition to walking state")
	j.animator.SetState(j.animator.WalkState)
}

type Walk struct {
	animator *Animator
}

func (w *Walk) playAnim() {
	fmt.Println("Playing walk animation")
}

func (w *Walk) transit() {
	fmt.Println("Transition to running state")
	w.animator.SetState(w.animator.RunState)
}
