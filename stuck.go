package infa

import (
	"fmt"
)

const TipExit = "type in 'exit' or 'quit'"

// Sometime, we need block the main goroutine, Otherwise the whole program will quit.
func BeBlock() {
	var cquit = make(chan bool)
	<-cquit
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
