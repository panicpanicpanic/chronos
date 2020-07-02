// Package cycle holds details about a 30 minute cycle
package cycle

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/panicpanicpanic/chronos/storage"
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
	var p Project
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

	p.Description, err = ProjectDescription.Ask()
	if err != nil {
		return c, err
	}

	p.JiraTicket, err = ProjectJiraTicket.Ask()
	if err != nil {
		return c, err
	}

	p.GithubPR, err = ProjectGithubPR.Ask()
	if err != nil {
		return c, err
	}

	c.Project = p

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

// RecapCurrentCycle launches a set of questions that aim to
// recap a current Cycle
func (c Cycle) RecapCurrentCycle(f *os.File) error {
	var r Recap
	var err error

	launchRecap, err := LaunchRecap.Ask()
	if err != nil {
		return err
	}
	if launchRecap == "Yes" {
		cycleCompleted, err := CycleCompleted.Ask()
		if err != nil {
			return err
		}
		if cycleCompleted == "Yes" {
			r.CycleCompleted = true
		} else {
			r.CycleCompleted = false
		}

		r.Distractions, err = CycleDistractions.Ask()
		if err != nil {
			return err
		}

		r.Improvements, err = CycleImprovements.Ask()
		if err != nil {
			return err
		}

		c.Recap = r
	}

	err = c.CycleToCSV(f)
	if err != nil {
		return err
	}

	return nil
}

// EndCycle asks the user if they want to stop working
func (c Cycle) EndCycle() (bool, error) {
	ec, err := EndCycle.Ask()
	if err != nil {
		fmt.Println(err)
	}
	if ec == "Yes" {
		return true, nil
	}

	return false, nil
}

func (c Cycle) launchRecap() {
	fmt.Println("Awesome! Here's a breakdown of this cycle:")
	fmt.Printf("Goal: %s \n", c.CycleGoal)
	fmt.Printf("Starting Point: %s \n", c.StartingPoint)
}

// CycleToCSV converts a Cycle to []string
// to write to CSV
func (c Cycle) CycleToCSV(f *os.File) error {

	data := []string{
		c.CycleGoal,
		c.CycleTitle,
		c.StartingPoint,
		c.Hazards,
		c.Energy,
		c.Morale,
		c.Duration.String(),
		c.StartTime.String(),
		c.EndTime.String(),
		strconv.FormatBool(c.Active),
		c.Project.Description,
		c.Project.JiraTicket,
		c.Project.GithubPR,
		strconv.FormatBool(c.Recap.CycleCompleted),
		c.Recap.Distractions,
		c.Recap.Improvements,
	}

	err := storage.Insert(f, data)
	if err != nil {
		return err
	}
	return nil
}
