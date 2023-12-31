package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/shirou/gopsutil/v3/process"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		text, _ := reader.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")
		var instructions = make([]IShite, 0)

		lines := strings.Split(text, " |")
		for _, subStr := range lines {
			subStr = strings.Trim(subStr, " ")
			if subStr == "" {
				continue
			}
			instructions = checkKWord(subStr, instructions)
		}

		for _, i := range instructions {
			i.doSomething()
		}
	}
}

func checkKWord(subStr string, instructions []IShite) []IShite {
	inputs := strings.Split(subStr, " ")
	switch inputs[0] {
	case "cd":
		return append(
			instructions,
			&CD{
				Instruction{
					input: inputs[1],
				},
			})
	case "pwd":
		return append(
			instructions,
			&PWD{},
		)
	case "echo":
		return append(
			instructions,
			&Echo{
				Instruction{
					input: inputs[1],
				},
			},
		)
	case "ps":
		return append(
			instructions,
			&PS{},
		)
	case "kill":
		return append(
			instructions,
			&Kill{
				Instruction{
					input: inputs[1],
				},
			},
		)
	case "fork":
		return append(
			instructions,
			&Fork{},
		)
	case "exec":
		return append(
			instructions,
			&Exec{
				Instruction{
					input: inputs[1],
				},
			},
		)
	default:
		return instructions
	}
}

type IShite interface {
	doSomething()
}

type Instruction struct {
	input string
}

func (i *Instruction) doSomething() {}

type CD struct {
	Instruction
}

func (i *CD) doSomething() {
	cd(i.input)
}

func cd(path string) {
	home, _ := os.UserHomeDir()
	err := os.Chdir(filepath.Join(home, path))

	if err != nil {
		err = os.Chdir(path)

		if err != nil {
			panic(err)
		}
	}

}

type PWD struct {
	Instruction
}

func (i *PWD) doSomething() {
	fmt.Println(pwd())
}

func pwd() string {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return path
}

type Echo struct {
	Instruction
}

func (i *Echo) doSomething() {
	echo(i.input)
}

func echo(msg string) {
	fmt.Println(msg)
}

type PS struct {
	Instruction
}

func (i *PS) doSomething() {
	for name, pid := range ps() {
		fmt.Println(name, pid)
	}
}

func ps() map[string]int {
	processes, _ := process.Processes()
	m := make(map[string]int, len(processes))
	for _, process := range processes {
		name, err := process.Name()

		if err == nil {
			m[name] = int(process.Pid)
		}
	}
	return m
}

type Kill struct {
	Instruction
}

func (i *Kill) doSomething() {
	pid, err := strconv.Atoi(i.input)

	if err != nil {
		return
	}

	kill(pid)
}

func kill(pid int) error {
	process, err := os.FindProcess(pid)

	if err != nil {
		return err
	}

	err = process.Kill()

	return err
}

type Fork struct {
	Instruction
}

func (i *Fork) doSomething() {
	fork()
}

func fork() {
	cmd := exec.Command("cmd", "/c", "start", os.Args[0])
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

type Exec struct {
	Instruction
}

func (i *Exec) doSomething() {
	myExec(i.input)
}

func myExec(exe string) {
	cmd := exec.Command(exe)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
