package pomodoro

import (
	"fmt"
	"time"
)

// Pomodoro Timer
// It is a
type Timer struct {
	delayTimeInSeconds int
	workTimeinMinutes  int
	breakTimeInMinutes int
}

// Construct a new timer
func NewTimer(delayTimeInSeconds int, workTimeinMinutes int, breakTimeInMinutes int) *Timer {
	return &Timer{delayTimeInSeconds: delayTimeInSeconds,
		workTimeinMinutes:  workTimeinMinutes,
		breakTimeInMinutes: breakTimeInMinutes}
}

func (t *Timer) Start(outputs chan<- string, stopSignal <-chan bool) {

	defer func() {
		close(outputs)
	}()
	// Delay time before starting
	var delayTimeInSeconds = t.delayTimeInSeconds
	for delayTimeInSeconds > 0 {
		outputs <- fmt.Sprintf("%03d...", delayTimeInSeconds)
		select {
		case <-time.After(time.Second):
			delayTimeInSeconds--
		}
	}

	sendBell(2, outputs)
	outputs <- fmt.Sprintf("Start working for %d minutes.", t.workTimeinMinutes)

	done := false
	inBreak := false

	for !done {
		if !inBreak {
			select {
			case <-time.After(time.Duration(t.workTimeinMinutes) * time.Minute):
				sendBell(2, outputs)
				outputs <- fmt.Sprintf("Take a break for %d minutes.", t.breakTimeInMinutes)
				inBreak = true
			case <-stopSignal:
				outputs <- "Bye bye.\nRemember to use Pomodoro the next time you want to be productive.\n"
				time.Sleep(2 * time.Second)
				done = true
			}
		} else {
			select {
			case <-time.After(time.Duration(t.breakTimeInMinutes) * time.Minute):
				sendBell(3, outputs)
				outputs <- fmt.Sprintf("Start working for %d minutes.", t.workTimeinMinutes)
				inBreak = false
			case <-stopSignal:
				outputs <- "Bye bye.\nRemember to use Pomodoro the next time you want to be productive.\n"
				time.Sleep(2 * time.Second)
				done = true
			}
		}
	}

}

func sendBell(times int, outputs chan<- string) {
	for i := 0; i < times; i++ {
		outputs <- "\a"
		time.Sleep(500 * time.Millisecond)
	}
}
