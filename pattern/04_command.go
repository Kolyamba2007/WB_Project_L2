package pattern

/*
Можно применить:

-Когда вы хотите параметризовать объекты выполняемым действием.
-Когда вы хотите ставить операции в очередь, выполнять их по расписанию или передавать по сети.
-Когда вам нужна операция отмены.

Плюсы:
-Убирает прямую зависимость между объектами, вызывающими операции,
и объектами, которые их непосредственно выполняют.
-Позволяет реализовать простую отмену и повтор операций.
-Позволяет реализовать отложенный запуск операций.
-Позволяет собирать сложные команды из простых.
-Реализует принцип открытости/закрытости.

Минусы:
-Усложняет код программы из-за введения множества дополнительных классов.

Ниже представлен пример видеоплеера - получателя, где кнопки - отправители.
Мы управляем функционалом плеера с помощью команд, которые отправляют кнопки.
При этом кнопки ничего не знают о самом функционале плеера.
*/

import "fmt"

type IInteractableView interface {
	Handle()
}

type Button struct {
	Command Command
}

func (b *Button) Handle() {
	b.Command.execute()
}

type Command interface {
	execute()
}

type PlayCommand struct {
	VideoPlayer VideoPlayer
}

func (c *PlayCommand) execute() {
	c.VideoPlayer.play()
}

type StopCommand struct {
	VideoPlayer VideoPlayer
}

func (c *StopCommand) execute() {
	c.VideoPlayer.stop()
}

type VideoPlayer interface {
	play()
	stop()
}

type LostfilmPlayer struct {
	isPlaying bool
}

func (p *LostfilmPlayer) play() {
	if !p.isPlaying {
		p.isPlaying = true
		fmt.Println("Playing the video player")
	}
}

func (p *LostfilmPlayer) stop() {
	if p.isPlaying {
		p.isPlaying = false
		fmt.Println("Stopping the video player")
	}
}
