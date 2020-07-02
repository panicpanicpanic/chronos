// Package cycle holds details about a 30 minute cycle
package cycle

import (
	"fmt"
	"os"
	"time"
)

// Cycle contains details about a Cycle of work
type Cycle struct {
	CycleGoal     string
	CycleTitle    string
	StartingPoint string
	Hazards       string
	Energy        string
	Morale        string
	Duration      time.Duration
	StartTime     time.Time
	EndTime       time.Time
	Active        bool
	Project
	Recap
}

// Project contains details about the project a Cycle
// is focusing on. This could be a particular JIRA ticket
// or a Github Pull Request
type Project struct {
	Description string
	JiraTicket  string
	GithubPR    string
}

// Recap contains details about recaps for a Cycle that
// has just ended
type Recap struct {
	CycleCompleted bool
	Distractions   string
	Improvements   string
}

// NewCycle returns a prepped Cycle for a user to work through
// It constructs a Cycle and returns details about the Cycle
// or returns an error
func NewCycle() (Cycle, error) {
	var c Cycle
	var err error

	c.CycleGoal, err = CycleAccomplishments.Ask()
	if err != nil {
		return c, err
	}

	c.StartingPoint, err = CycleGetStarted.Ask()
	if err != nil {
		return c, err
	}

	c.Hazards, err = CycleHazrds.Ask()
	if err != nil {
		return c, err
	}

	c.Energy, err = CycleEnergy.Ask()
	if err != nil {
		return c, err
	}

	c.Morale, err = CycleMorale.Ask()
	if err != nil {
		return c, err
	}

	duration, err := CycleDuration.Ask()
	if err != nil {
		return c, err
	}

	c.Duration, err = time.ParseDuration(duration)
	if err != nil {
		return c, err
	}

	c.launchRecap()

	launch, err := LaunchCycle.Ask()
	if err != nil {
		return c, err
	}
	if launch == "No" {
		fmt.Println("Okay...going to leave now!")
		os.Exit(1)
	}

	return c, nil
}

// RecapCurrentCycle
func (c Cycle) RecapCurrentCycle() error {

	return nil
}

func (c Cycle) launchRecap() {
	fmt.Println("Awesome! Here's a breakdown of this cycle:")
	fmt.Printf("Goal:%s \n", c.CycleGoal)
	fmt.Printf("Starting Point:%s \n", c.StartingPoint)
}
