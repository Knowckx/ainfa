package infa

import (
	"fmt"
)

const TipExit = "type in 'exit' or 'quit'"

func BeStuck() {
	fmt.Println("wait for cancel sign. " + TipExit)
	var cquit = make(chan bool)
	go GetInput(cquit)
	<-cquit
	fmt.Println("bye-bye")
}

func Stuck() {
	var c = make(chan bool)
	<-c
}

func GetInput(c chan<- bool) {
	input := ""
	maxLoop := 3
	cnt := 0
	for {
		cnt++
		fmt.Scanln(&input)
		if input == "exit" || input == "quit" {
			break
		}
		fmt.Println("input is", input)
		if cnt <= maxLoop {
			fmt.Println(TipExit)
		} else {
			break
		}
	}
	c <- true
}
