package main

import (
	"fmt"
	"os"
	"time"

	"github.com/panicpanicpanic/chronos/cycle"
	"github.com/panicpanicpanic/chronos/storage"
)

func main() {
	// Create CSV if doesn't already exist
	f, err := storage.WriteCSV()
	if err != nil {
		fmt.Println(err)
	}

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

		err = cycle.RecapCurrentCycle(f)
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
