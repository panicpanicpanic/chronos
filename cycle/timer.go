package cycle

import (
	"time"

	"github.com/drgrib/ttimer/agent"
)

// StartTimer creates a new timer and starts a countdown
// with details about the duration coming on the Cycle
func (c *Cycle) StartTimer() error {
	c.StartTime = time.Now()

	t := agent.Timer{
		Title: fmt.Sprintf("Goal:%s | StartPoint: %s", c.CycleGoal, c.StartingPoint)
	}
	c.EndTime = time.Now()

	t.Start(c.Duration)
	t.CountDown()

	return nil
}
