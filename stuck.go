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

func GetInput(c chan<- bool) {
	input := ""
	for {
		fmt.Scanln(&input)
		if input == "exit" || input == "quit" {
			break
		}
		fmt.Println(TipExit)
	}
	c <- true
}
