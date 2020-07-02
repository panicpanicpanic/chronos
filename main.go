package main

import (
	"fmt"
	"time"

	"github.com/panicpanicpanic/tomato/cycle"
)

func main() {
	var err error

	cycle, err := cycle.NewCycle()
	if err != nil {
		fmt.Println(err)
	}

	err = cycle.StartTimer()
	if err != nil {
		fmt.Println(err)
	}

	// This line will hit once the timer ends or quits
	cycle.EndTime = time.Now()

	err = cycle.RecapCurrentCycle()
	if err != nil {
		fmt.Println(err)
	}

	// if the timer is disrupted:
	// if the countdown channel is closed, ask if they want to recap (bool)
	// if no, exit

	// if yes, start the recap chat
	// recaping cycle
	// Completed cycle's target?
	// Anything noteworthy?
	// Any distractions?
	// Things to improve for next cycle?
	// start a new cycle
}
