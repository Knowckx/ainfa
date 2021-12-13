package infa

import (
	"fmt"
	"log"
	"time"
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

// RecordExecTime Record function execution time
// useage: defer ExecTime(functionName)()
func RecordExecTime(funcName string) func() {
	start := time.Now()
	log.Printf("enter %s", funcName)
	return func() {
		log.Printf("exit. Time Usege:%s", time.Since(start))
	}
}
