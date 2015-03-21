package main

import (
	"fmt"
	"github.com/erezak/pomodoro/pomodoro"
)

func main() {
	timerOutputs := make(chan string, 1)
	stopSingal := make(chan bool)

	var timer = pomodoro.NewTimer(5, 25, 5)
	fmt.Println("Starting timer - press enter to stop.")
	go timer.Start(timerOutputs, stopSingal)

	go func() {
		fmt.Scanln()
		stopSingal <- true
	}()

	for output := range timerOutputs {
		fmt.Println(output)
	}

}
