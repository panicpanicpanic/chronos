package main

import (
	"fmt"
	"os"
	"time"

	"github.com/panicpanicpanic/chronos/cycle"
)

func main() {
	var cycleLoop func()
	// Kickoff cycle
	cycleLoop = func() {
		var err error

		cycle, err := cycle.NewCycle()
		if err != nil {
			fmt.Println(err)
		}

		err = cycle.StartTimer()
		if err != nil {
			fmt.Println(err)
		}

		cycle.EndTime = time.Now()

		err = cycle.RecapCurrentCycle()
		if err != nil {
			fmt.Println(err)
		}

		end, err := cycle.EndCycle()
		if err != nil {
			fmt.Println(err)
		}

		if end {
			os.Exit(1)
		} else {
			cycleLoop()
		}
	}
	cycleLoop()
}
