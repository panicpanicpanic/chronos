package cycle

import (
	"errors"

	"github.com/manifoldco/promptui"
)

// Prompt holds references to the prompt types for a Cycle
type Prompt struct {
	Question *promptui.Prompt
	DropDown *promptui.Select
}

// Ask prompts a Question or Dropdown picker
// It returns the value a User inputs, or an error
func (p Prompt) Ask() (string, error) {
	var err error
	var result string

	switch {
	case p.Question != nil:
		result, err = p.Question.Run()
		if err != nil {
			return "", err
		}
	case p.DropDown != nil:
		_, result, err = p.DropDown.Run()
		if err != nil {
			return "", err
		}
	}

	return result, nil
}

/* Questions */

// CycleAccomplishments - What am I trying to accomplish this cycle?
var CycleAccomplishments = Prompt{
	Question: &promptui.Prompt{
		Label: "What am I trying to accomplish this cycle?",
		Validate: func(input string) error {
			if input == "" {
				return errors.New("cant be empty")
			}
			return nil
		},
	},
}

// CycleGetStarted - How will I get started?
var CycleGetStarted = Prompt{
	Question: &promptui.Prompt{
		Label: "How will I get started?",
		Validate: func(input string) error {
			if input == "" {
				return errors.New("cant be empty")
			}
			return nil
		},
	},
}

// CycleHazrds - Any hazards present?
var CycleHazrds = Prompt{
	Question: &promptui.Prompt{
		Label: "Any hazards present?",
		Validate: func(input string) error {
			if input == "" {
				return errors.New("cant be empty")
			}
			return nil
		},
	},
}

// CycleEnergy - What is your current energy level?
var CycleEnergy = Prompt{
	DropDown: &promptui.Select{
		Label: "What is your current energy level?",
		Items: []string{"High", "Medium", "Low"},
	},
}

// CycleMorale - What is your current morale level?
var CycleMorale = Prompt{
	DropDown: &promptui.Select{
		Label: "What is your current morale level?",
		Items: []string{"High", "Medium", "Low"},
	},
}

// CycleDuration - How long should this Cycle be?
var CycleDuration = Prompt{
	DropDown: &promptui.Select{
		Label: "How long should this Cycle be (in minutes)?",
		Items: []string{"15m", "30m", "45m"},
	},
}

// LaunchCycle - Ready to start this cycle?
var LaunchCycle = Prompt{
	DropDown: &promptui.Select{
		Label: "Ready to start this cycle?",
		Items: []string{"Yes", "No"},
	},
}

// LaunchRecap - Did you want to recap this cycle?
var LaunchRecap = Prompt{
	DropDown: &promptui.Select{
		Label: "Did you want to recap this cycle?",
		Items: []string{"Yes", "No"},
	},
}

// CycleCompleted - Did you complete the cycle's target?
var CycleCompleted = Prompt{
	DropDown: &promptui.Select{
		Label: "Did you complete the cycle's target?",
		Items: []string{"Yes", "No"},
	},
}

// CycleDistractions - Were there any distractions?
var CycleDistractions = Prompt{
	Question: &promptui.Prompt{
		Label: "Were there any distractions?",
		Validate: func(input string) error {
			if input == "" {
				return errors.New("cant be empty")
			}
			return nil
		},
	},
}

// CycleImprovements - Are there any improvements you would make?
var CycleImprovements = Prompt{
	Question: &promptui.Prompt{
		Label: "Are there any improvements you would make?",
		Validate: func(input string) error {
			if input == "" {
				return errors.New("cant be empty")
			}
			return nil
		},
	},
}

// ProjectDescription - Quick description of the project you're working on this cycle
var ProjectDescription = Prompt{
	Question: &promptui.Prompt{
		Label: "Quick description of the project you're working on this cycle",
		Validate: func(input string) error {
			if input == "" {
				return errors.New("cant be empty")
			}
			return nil
		},
	},
}

// ProjectJiraTicket - Any JIRA ticket/links you're going to focus on this cycle
var ProjectJiraTicket = Prompt{
	Question: &promptui.Prompt{
		Label: "Any JIRA ticket/links you're going to focus on this cycle",
		Validate: func(input string) error {
			if input == "" {
				return errors.New("cant be empty")
			}
			return nil
		},
	},
}

// ProjectGithubPR - Any Github PR/Issue you're going to focus on this cycle
var ProjectGithubPR = Prompt{
	Question: &promptui.Prompt{
		Label: "Any Github PR/Issue you're going to focus on this cycle",
		Validate: func(input string) error {
			if input == "" {
				return errors.New("cant be empty")
			}
			return nil
		},
	},
}

// EndCycle - Did you want to stop working?
var EndCycle = Prompt{
	DropDown: &promptui.Select{
		Label: "Did you want to stop working?",
		Items: []string{"Yes", "No"},
	},
}
